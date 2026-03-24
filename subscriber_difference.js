// Ручное сравнение подписчиков github

// Собери со всех страниц
function usersList() {
	const followersData = Array.from(
		document.querySelectorAll("#user-profile-frame .Link--secondary, .follow-list-item .Link--secondary")
	).map(el => el.textContent.trim()).join("\n")
	console.log(followersData)
	
	return followersData
}


/**
 * Сравнивает списки подписчиков и показывает, кто не подписан в ответ
 * @param {string} followers - список подписчиков (через \n)
 * @param {string} following - список тех, на кого вы подписаны (через \n)
 */
function findNotFollowingBack(followers, following) {
    // Парсим строки в массивы, убираем пустые строки и лишние пробелы
    const followersSet = new Set(
        followers.split("\n")
            .map(s => s.trim())
            .filter(Boolean)
    );
    
    const followingList = following.split("\n")
            .map(s => s.trim())
            .filter(Boolean);
    
    // Находим тех, на кого вы подписаны, но кто не подписан на вас
    const notFollowingBack = followingList.filter(user => !followersSet.has(user));
    
    // Выводим результат в консоль
    console.log(`🔍 Найдено ${notFollowingBack.length} пользователей, которые не подписаны на вас:`);
    notFollowingBack.forEach((user, i) => {
        console.log(`${i + 1}. https://github.com/${user}`);
    });
    
    // Возвращаем массив для дальнейшего использования (опционально)
    return notFollowingBack;
}


findNotFollowingBack(``, ``)

