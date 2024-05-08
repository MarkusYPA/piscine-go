package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	bn := 48 // base number for printing digits

	for i := 0; i < 10; i++ {
		if n > 1 {
			for j := i + 1; j < 10; j++ {
				if n > 2 {
					for k := j + 1; k < 10; k++ {
						if n > 3 {
							for l := k + 1; l < 10; l++ {
								if n > 4 {
									for m := l + 1; m < 10; m++ {
										if n > 5 {
											for nn := m + 1; nn < 10; nn++ {
												if n > 6 {
													for o := nn + 1; o < 10; o++ {
														if n > 7 {
															for p := o + 1; p < 10; p++ {
																if n > 8 {
																	for q := p + 1; q < 10; q++ {

																		z01.PrintRune(rune(i + bn))
																		z01.PrintRune(rune(j + bn))
																		z01.PrintRune(rune(k + bn))
																		z01.PrintRune(rune(l + bn))
																		z01.PrintRune(rune(m + bn))
																		z01.PrintRune(rune(nn + bn))
																		z01.PrintRune(rune(o + bn))
																		z01.PrintRune(rune(p + bn))
																		z01.PrintRune(rune(q + bn))

																		if i < 1 {
																			z01.PrintRune(',')
																			z01.PrintRune(' ')
																		}
																	}
																} else {

																	z01.PrintRune(rune(i + bn))
																	z01.PrintRune(rune(j + bn))
																	z01.PrintRune(rune(k + bn))
																	z01.PrintRune(rune(l + bn))
																	z01.PrintRune(rune(m + bn))
																	z01.PrintRune(rune(nn + bn))
																	z01.PrintRune(rune(o + bn))
																	z01.PrintRune(rune(p + bn))

																	if i < 2 {
																		z01.PrintRune(',')
																		z01.PrintRune(' ')
																	}
																}
															}
														} else {

															z01.PrintRune(rune(i + bn))
															z01.PrintRune(rune(j + bn))
															z01.PrintRune(rune(k + bn))
															z01.PrintRune(rune(l + bn))
															z01.PrintRune(rune(m + bn))
															z01.PrintRune(rune(nn + bn))
															z01.PrintRune(rune(o + bn))

															if i < 3 {
																z01.PrintRune(',')
																z01.PrintRune(' ')
															}
														}
													}
												} else {

													z01.PrintRune(rune(i + bn))
													z01.PrintRune(rune(j + bn))
													z01.PrintRune(rune(k + bn))
													z01.PrintRune(rune(l + bn))
													z01.PrintRune(rune(m + bn))
													z01.PrintRune(rune(nn + bn))

													if i < 4 {
														z01.PrintRune(',')
														z01.PrintRune(' ')
													}
												}
											}
										} else {

											z01.PrintRune(rune(i + bn))
											z01.PrintRune(rune(j + bn))
											z01.PrintRune(rune(k + bn))
											z01.PrintRune(rune(l + bn))
											z01.PrintRune(rune(m + bn))

											if i < 5 {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								} else {

									z01.PrintRune(rune(i + bn))
									z01.PrintRune(rune(j + bn))
									z01.PrintRune(rune(k + bn))
									z01.PrintRune(rune(l + bn))

									if i < 6 {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						} else {

							z01.PrintRune(rune(i + bn))
							z01.PrintRune(rune(j + bn))
							z01.PrintRune(rune(k + bn))

							if i < 7 {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				} else {
					z01.PrintRune(rune(i + bn))
					z01.PrintRune(rune(j + bn))

					if i < 8 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		} else {
			z01.PrintRune(rune(i + bn))

			if i < 9 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}

	z01.PrintRune('\n')
}
