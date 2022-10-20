all: clean app

clean:
	rm -f browurls

app:
	go build -o browurls ./cmd