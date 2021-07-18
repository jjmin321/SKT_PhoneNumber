# SKT_PhoneNumber
좋은 전화번호를 찾아 변경하기 위해 만든 프로그램입니다.

config 폴더에 config.toml을 생성한 뒤 config.go에 있는 skt 구조체의 각 필드에 맞는 값을 아래와 같이 설정하고, line_num_from과 line_num_to를 변경하여 원하는 전화번호를 조회할 수 있습니다.

# config.toml 예시
<img width="338" alt="image" src="https://user-images.githubusercontent.com/52072077/126069807-d6383187-3012-4ef5-b78b-91ec02e6eda2.png">

# 기능 추가 및 응용
기능을 추가하고 싶다면 아래와 같은 코드를 추가하여 사용하면 됩니다.

또한 이를 응용하여 프로그램이 자동으로 전화번호를 조회하며, 알고리즘이 좋은 전화번호를 찾았을 경우 사용자의 휴대폰으로 알림을 보내는 등의 기능을 추가할 수 있습니다.

```go
func parseMainNodes(n *html.Node) bool {
	if n.DataAtom == atom.Strong {
		return scrape.Attr(n, "class") == "phoneNumber"
	}
	return false
}

root, err := html.Parse(resp.Body)

phoneNumberList := scrape.FindAll(root, parseMainNodes)

for _, phoneNumber := range phoneNumberList {
	log.Info.Println(scrape.Text(phoneNumber))
}
```