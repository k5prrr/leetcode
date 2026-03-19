package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type ICheckSites interface {
	Add(url string) error
	AddList(urls []string) error
	Delete(url string) error
	Clear() error
	Check(ctx context.Context) map[string]CheckSitesAnswerItem
	Show(ctx context.Context) string
}
type CheckSitesAnswerItem struct {
	URL  string
	Code int
	Data string
}
type CheckSites struct {
	Sites    []string
	MapSites map[string]interface{}
	PathFile string
	mu       sync.Mutex
}

func NewCheckSites(pathFile string) ICheckSites {
	if pathFile == "" {
		pathFile = "sites.txt"
	}
	cs := &CheckSites{
		PathFile: pathFile,
		MapSites: make(map[string]interface{}),
	}
	cs.loadList()

	return cs
}
func (c *CheckSites) loadList() {
	if _, err := os.Stat(c.PathFile); err != nil {
		c.Sites = make([]string, 0)
		return
	}

	file, err := os.Open(c.PathFile)
	if err != nil {
		c.Sites = make([]string, 0)
		return
	}
	defer file.Close()

	var sites []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		site := strings.TrimSpace(scanner.Text())
		if site != "" {
			sites = append(sites, site)
		}
	}

	if err := scanner.Err(); err != nil {
		c.Sites = make([]string, 0)
		return
	}

	// Заполнение мапы и слайса
	//c.mu.Lock()
	//defer c.mu.Unlock()

	for _, site := range sites {
		c.MapSites[site] = struct{}{}
	}
	sort.Strings(sites)
	c.Sites = sites
}
func (c *CheckSites) saveList() error {
	if len(c.Sites) == 0 {
		err := os.Remove(c.PathFile)
		if err != nil && !os.IsNotExist(err) {
			log.Printf("Error removing file %s: %v", c.PathFile, err)
		}
		return nil
	}

	text := strings.Join(c.Sites, "\n")

	file, err := os.OpenFile(c.PathFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	return err
}

func (c *CheckSites) Add(url string) error {
	url = strings.TrimSpace(url)
	if url == "" {
		return errors.New("empty URL provided")
	}
	url = c.useProtocolScheme(url)

	if _, have := c.MapSites[url]; have {
		return errors.New("URL already exists")
	}

	c.MapSites[url] = struct{}{}
	c.Sites = append(c.Sites, url)
	sort.Strings(c.Sites)

	return c.saveList()
}
func (c *CheckSites) AddList(urls []string) error {
	var mu sync.Mutex
	newSites := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(rawURL string) {
			defer wg.Done()

			url := strings.TrimSpace(rawURL)
			if url == "" {
				return
			}
			url = c.useProtocolScheme(url)

			// Локальная проверка — без блокировки
			c.mu.Lock()
			_, exists := c.MapSites[url]
			c.mu.Unlock()

			if exists {
				return
			}

			// Собираем новые в локальную мапу
			mu.Lock()
			newSites[url] = struct{}{}
			mu.Unlock()
		}(url)
	}
	wg.Wait()

	// Атомарно добавляем всё новое
	if len(newSites) > 0 {
		c.mu.Lock()
		for url := range newSites {
			if _, exists := c.MapSites[url]; !exists {
				c.MapSites[url] = struct{}{}
				c.Sites = append(c.Sites, url)
			}
		}
		sort.Strings(c.Sites)
		c.mu.Unlock()
	}

	return c.saveList()
}
func (c *CheckSites) Delete(url string) error {
	url = strings.TrimSpace(url)

	if _, have := c.MapSites[url]; !have {
		return errors.New("URL not found")
	}

	delete(c.MapSites, url)

	for i := range c.Sites {
		if c.Sites[i] == url {
			c.Sites = append(c.Sites[:i], c.Sites[i+1:]...)
			break
		}
	}

	return c.saveList()
}
func (c *CheckSites) Clear() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Sites = make([]string, 0)
	c.MapSites = make(map[string]interface{})

	return os.Remove(c.PathFile)
}
func (c *CheckSites) useProtocolScheme(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	return url
}
func (c *CheckSites) Check(ctx context.Context) map[string]CheckSitesAnswerItem {
	//result := make(map[string]int)
	resultChan := make(chan CheckSitesAnswerItem, len(c.Sites))

	c.mu.Lock()
	sitesCopy := make([]string, len(c.Sites))
	copy(sitesCopy, c.Sites)
	c.mu.Unlock()

	wg := sync.WaitGroup{}
	wg.Add(len(c.Sites))

	for _, site := range c.Sites {
		go func(site string) {
			defer wg.Done()

			item := CheckSitesAnswerItem{
				URL: site,
			}

			status, location, err := c.StatusCode(ctx, site)
			if err != nil {
				item.Code = 0
				item.Data = err.Error()
			} else {
				item.Code = status
				item.Data = location
				if status >= 200 && status < 300 {
					item.Data = "OK"
				}
			}

			resultChan <- item
		}(site)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	result := make(map[string]CheckSitesAnswerItem)
	for item := range resultChan {
		result[item.URL] = item
	}

	return result
}
func (c *CheckSites) Show(ctx context.Context) string {
	var result strings.Builder
	mapResult := c.Check(ctx)
	for i, site := range c.Sites {
		result.WriteString(fmt.Sprintf("%d. %s ", i+1, site))
		if item, ok := mapResult[site]; ok {
			if item.Code == 200 {
				result.WriteString("✅")
			} else {
				result.WriteString(fmt.Sprintf("❌ %d %s", item.Code, item.Data))
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}

func (c *CheckSites) StatusCode(ctx context.Context, url string) (int, string, error) {
	client := &http.Client{
		Timeout: 4 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	location := ""
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location = resp.Header.Get("Location")
	}

	return resp.StatusCode, location, nil
}

// Use
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 16*time.Second)
	defer cancel()

	checkSites := NewCheckSites("")
	checkSites.AddList([]string{
		"h7team.ru",
		"pbrain.ru",
		"partbrain.ru",
		"app.partbrain.ru",
		"tvpaty.ru",
		"tvparty.ru",
		"st-autoservice.ru",
		"nejnosex.ru",
		"gway-logistic.ru",
		"velinbrand.com",
	})
	checkSites.AddList([]string{
		"armango.com",
		"test.armango.com",
		"report.armango.com",

		"brusko.ru",
		"academy.brusko.ru",
		"bot.brusko.ru",

		"seller.brusko.ru",
		"partners.brusko.ru",
		"fabrikabrusko.ru",

		"b24.brusko.ru",
		"vapenews.ru",
		"awards.vapenews.ru",

		"angryvape.ru",
		"sexologic.ru",
		"pro-baits.pro",

		"happman.ru",
		"innoheat.ru",
		"nejnosex.ru",

		"botya.brusko.ru",

		"azoral.ru",
		"2x2dent.ru",
		"analit.pro",
	})

	fmt.Println(checkSites.Show(ctx))
}
