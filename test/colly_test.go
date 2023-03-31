package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCheckGames(t *testing.T) {
	//52 14 60 7 83 67 15 64 4 98 87 81 97 10 57 28 82 85 45 80 1 99 31 79 84 86 22 88 102 101 106 112
	//115 123 129 149 151 153 156 191 195 217 227 158 163 161 179 188 200 172 175 259 264 265 270 271 272
	//273 295 299 296 297 298 308 278 333 334 335 240 238 348 351 332 354 356 320 377 378 379 282 288 294
	//331 381 256 329 391 398 303 402 367 330 372 307 409 410 411 397 321 389 420 390 423 430 413 380 443
	//446 447 370 394 437 292 357 461 369 462 371 445 464 465 431 466 342 467 468 469 474 475 386 442 418
	//451 426 454 488 417 490 491 440 495 505 476 416 428 405 463 518 522 448 523 434 449 450 480 453 427
	//525 441 531 521 506 533 473 537 484 538 470 487 407 520 500 502 492 419 526 436 496 400 528 501 456
	//554 511 512 507 510 513 412 444 566 527 485 570 373 572 573 574 576 577 432 404 582 586 588 534 547
	//550 619 620 626 627 628 532 535 542 631 652 653 622 654 655 569 632 656 666 555 667 668 559 670 557
	//624 551 623 585 556 629 630 625 515 517 258 514 499 516 529 433 932 933 935 936 937 938 934 1144
	//1145 1146 1147 1148 1149 1156 1158 1170 1171 1172 1173 1180 1182 1092 1164 1193 1194 1196 1152
	//1155 1174 1198 1151 1199 1200 1203 1204 1205 1169 1206 1192 1207 1090 1208 1209 1175 1179 1150
	//1091 1181 1154 1212 1214 1202 1166 1094 1211 1227 1177 1153 1176 1239 1242 1244 1165 1245 1178
	//1216 1217 1229 1254 1201 1257 1232 1234 1238 1219 1248 1189 1226 1225 1188 1167 1190 1195 1213
	//1163 1210 1222 1243 1168 1255 1253 1218 1223 1221 1231 1237 1256 1240 1224 1197 1258 1191 1228
	//1215 1246 1230 1220 1236 1247 1233 1235 1261 1260 1262 1093 1259
	running := 0
	hasGameNumerics := make([]int, 0)
	numerics := make(chan int)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case k := <-numerics:
				hasGameNumerics = append(hasGameNumerics, k)
			case <-done:
				running--
			}
			if running == 0 {
				break
			}
		}
	}()

	i := 0
	for {
		if running <= 100 {
			go func(p int) {
				if g, err := http.Get(fmt.Sprintf("https://static.game.fm/data/%d/index.html", p)); err == nil {
					if g.StatusCode == 200 {
						numerics <- p
					}
					done <- struct{}{}
				}
			}(i)
			running++
			i++
		}

		if i > 2000 {
			break
		}
	}

	fmt.Println(len(hasGameNumerics), hasGameNumerics)
}

func TestColly(t *testing.T) {
	// https://static.game.fm/data/1/index.html
	//c := colly.NewCollector()
	//
	//c.OnError(func(res *colly.Response, err error) {
	//	fmt.Println(err)
	//})
	//
	//c.OnResponse(func(res *colly.Response) {
	//	//fmt.Println(string(res.Body))
	//})
	//c.OnHTML("script", func(e *colly.HTMLElement) {
	//	fmt.Println(e.Attr("src"))
	//})
	//c.OnHTML("link", func(e *colly.HTMLElement) {
	//	fmt.Println(e.Attr("href"))
	//})
	//
	//_ = c.Visit("https://static.game.fm/data/1/index.html")

	n := 1253
	for n > 0 {
		fmt.Println(n%10, n)
		n /= 10
	}
}
