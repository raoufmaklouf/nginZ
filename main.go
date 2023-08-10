package main

import (
	"bufio"
	"os"
	"sync"
)

var client, tr = createHTTPClient()

func main() {

	defer client.CloseIdleConnections()
	var wg sync.WaitGroup
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {

		line := sc.Text()

		if isUrl(line) == true {
			cleanurl := line
			lastchar := line[len(line)-1:]
			if lastchar == "/" {
				cleanurl = line[:len(line)-1]
			}
			wg.Add(7)
			go func() {
				defer wg.Done()
				OffBySlash(cleanurl)
			}()
			go func() {
				defer wg.Done()
				SCRIPT_NAME(cleanurl)
			}()
			go func() {
				defer wg.Done()
				uriCRLF(cleanurl)
			}()
			go func() {
				defer wg.Done()
				AnyVariable(cleanurl)
			}()
			go func() {
				defer wg.Done()
				HttpRequestSplitting(cleanurl)
			}()
			go func() {
				defer wg.Done()
				controllingSocket(cleanurl)
			}()
			go func() {
				defer wg.Done()
				Xcrlf(cleanurl)
			}()

		}

	}
	wg.Wait()

}
