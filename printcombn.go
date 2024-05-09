package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	bn := 48 // base number for printing digits

	if n == 1 {
		for i := 0; i < 10; i++ {

			z01.PrintRune(rune(i + bn))

			if i < 9 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}

	if n == 2 {
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {

				z01.PrintRune(rune(i + bn))
				z01.PrintRune(rune(j + bn))

				if i < 8 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}

	if n == 3 {
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {

					z01.PrintRune(rune(i + bn))
					z01.PrintRune(rune(j + bn))
					z01.PrintRune(rune(k + bn))

					if i < 7 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}

	if n == 4 {
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {

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
			}
		}
	}

	if n == 5 {
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {

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
				}
			}
		}
	}

	if n == 6 {
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							for nn := m + 1; nn < 10; nn++ {

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
					}
				}
			}
		}
	}

	if n == 7 {
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							for nn := m + 1; nn < 10; nn++ {
								for o := nn + 1; o < 10; o++ {

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
						}
					}
				}
			}
		}
	}

	if n == 8 {
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							for nn := m + 1; nn < 10; nn++ {
								for o := nn + 1; o < 10; o++ {
									for p := o + 1; p < 10; p++ {

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
							}
						}
					}
				}
			}
		}
	}

	if n == 9 {
		for i := 0; i < 10; i++ {
			for j := i + 1; j < 10; j++ {
				for k := j + 1; k < 10; k++ {
					for l := k + 1; l < 10; l++ {
						for m := l + 1; m < 10; m++ {
							for nn := m + 1; nn < 10; nn++ {
								for o := nn + 1; o < 10; o++ {
									for p := o + 1; p < 10; p++ {
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
									}
								}
							}
						}
					}
				}
			}
		}
	}

	z01.PrintRune('\n')
}
