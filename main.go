/*
 * File: main.go
 * Created: Sunday, 7th February 2021 4:31:07 pm
 * Author: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Last Modified: Friday, 12th February 2021 4:12:46 am
 * Modified By: Rakibul Yeasin (ryeasin03@gmail.com)
 * -----
 * Copyright (c) 2021 Slishee
 */

package main

import (
	"time"
	driver "webdriver/webdriver"
)

func main() {
	url := `https://google.com`
	driver.RunServer()
	driver.GetSession()
	driver.Get(url)
	time.Sleep(8 * time.Second)
	driver.Screenshot("google")
	time.Sleep(8 * time.Second)

	defer driver.Kill()
}
