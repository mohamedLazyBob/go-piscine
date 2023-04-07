package piscine

func ReverseMenuIndex(menu []string) []string {
	revMenu := make([]string, len(menu))
	for i, j := 0, len(menu)-1; i < len(revMenu); i, j = i+1, j-1 {
		revMenu[i] = menu[j]
	}
	return revMenu
}
