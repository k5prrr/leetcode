#!/bin/bash
#
# Скрипт для анализа односторонних подписок на GitHub (для Ubuntu)
# Выводит два списка:
#   1. Вы подписаны на пользователя, но он на вас — нет
#   2. Пользователь подписан на вас, но вы на него — нет
#
# Использование: 
#   chmod +x git_checking_mutual_subscription.sh 
#   ./git_checking_mutual_subscription.sh

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info()    { echo -e "${GREEN}[INFO]${NC} $1"; }
log_warn()    { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error()   { echo -e "${RED}[ERROR]${NC} $1" >&2; }
log_header()  { echo -e "\n${BLUE}═══ $1 ═══${NC}\n"; }

# ─────────────────────────────────────────────────────────────
# 1. Проверка и установка gh
# ─────────────────────────────────────────────────────────────
if ! command -v gh &> /dev/null; then
    log_warn "GitHub CLI (gh) не установлен. Устанавливаю..."
    echo "Требуется ввод пароля sudo для установки пакетов."
    sudo apt update
    sudo apt install -y gh
    log_info "gh успешно установлен."
else
    log_info "gh уже установлен: $(gh --version | head -n1)"
fi

# ─────────────────────────────────────────────────────────────
# 2. Проверка авторизации
# ─────────────────────────────────────────────────────────────
if ! gh auth status &> /dev/null; then
    log_warn "Вы не авторизованы в GitHub CLI."
    log_info "Запускаю процесс авторизации..."
    echo ""
    echo ">>> Следуйте инструкциям в браузере."
    echo ">>> После успешного входа вернитесь в терминал и нажмите Enter."
    echo ""
    gh auth login
    log_info "Авторизация завершена!"
else
    log_info "Вы авторизованы!"
fi

# ─────────────────────────────────────────────────────────────
# 3. Подготовка данных (кешируем списки)
# ─────────────────────────────────────────────────────────────
echo ""
log_info "Получаю данные от GitHub API (это может занять время)..."

FOLLOWERS_CACHE=$(mktemp)
FOLLOWING_CACHE=$(mktemp)

fetch_with_retry() {
    local endpoint=$1
    local output=$2
    local attempt=0
    local max_attempts=3
    
    while [ $attempt -lt $max_attempts ]; do
        if gh api "$endpoint" --paginate --jq '.[].login' 2>/dev/null | sort > "$output"; then
            if [ -s "$output" ]; then
                return 0
            fi
        fi
        attempt=$((attempt + 1))
        log_warn "Попытка $attempt/$max_attempts не удалась, повторяю..."
        sleep 2
    done
    return 1
}

if ! fetch_with_retry "user/followers" "$FOLLOWERS_CACHE"; then
    log_error "Не удалось получить список подписчиков."
    rm -f "$FOLLOWERS_CACHE" "$FOLLOWING_CACHE"
    exit 1
fi

if ! fetch_with_retry "user/following" "$FOLLOWING_CACHE"; then
    log_error "Не удалось получить список подписок."
    rm -f "$FOLLOWERS_CACHE" "$FOLLOWING_CACHE"
    exit 1
fi

log_info "Данные получены. Формирую списки..."

# ─────────────────────────────────────────────────────────────
# Список 1: Вы подписаны, но на вас не подписаны
# ─────────────────────────────────────────────────────────────
log_header "ВЫ подписаны на них, но они на вас - НЕТ"

comm -13 "$FOLLOWERS_CACHE" "$FOLLOWING_CACHE" | \
    sed 's|^|https://github.com/|' | \
    cat -n

# ─────────────────────────────────────────────────────────────
# Список 2: На вас подписаны, но вы не подписаны
# ─────────────────────────────────────────────────────────────
log_header "На вас подписаны, но ВЫ на них - НЕТ"

comm -23 "$FOLLOWERS_CACHE" "$FOLLOWING_CACHE" | \
    sed 's|^|https://github.com/|' | \
    cat -n

# ─────────────────────────────────────────────────────────────
# Очистка и завершение
# ─────────────────────────────────────────────────────────────
rm -f "$FOLLOWERS_CACHE" "$FOLLOWING_CACHE"

echo ""
log_info "Готово!"
