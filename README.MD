<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/iskaa02/qalam">
    <h1>✏️</h1>
  </a>

<h2 align="center">Qalam</h2>
  <p align="center">
    Easy Terminal styling using go
    <br />
    <br />
    <a href="https://github.com/iskaa02/qalam/issues">Report Bug</a>
    ·
    <a href="https://github.com/iskaa02/qalam/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li><a href="#styler">Using Styler</a></li>
        <li><a href="#bbcode">Using BBcode</a></li>
      </ul>
    </li>
    <li><a href="#bbcode-examples">BBcode Exmaples</a></li>
    <li><a href="#gradient">Gradient</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#notes">Notes</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

[![Product Name Screen Shot][product-screenshot]](./doc/images/qalam_example.png)

Qalam is A terminal Styling library built for Go to make building cool-looking CLIs easier

<p align="right">(<a href="#top">back to top</a>)</p>

## Getting Started

There is Two Ways to use Qalam

1. Using [Styler](#with-styler)
2. Using [BBcode format](#with-bbcode-format)

### Installation

```bash
go get github.com/iskaa02/qalam
```

## Usage

### With Styler

```go
import "github.com/iskaa02/qalam"
func main(){
	s:=qalam.NewStyler().
		Bold().
		Italic().
		Red()
	// This will print bold italic red text to the terminal
	s.Print("Hello world")

	// Change background color
	s.YellowBackground()

	// use hex color as foreground color
	s.Hex("#e28421")

	// use hex color as background color
	s.BackgroundHex("#e28421")

	// Flip Foreground and background colors
	s.Invert()

	// Make Text blink
	s.Blink()

}
```

Removing Styles

```go
// You can remove any style added previously
 s.Remove("bold","blink","bg#e28421")
```

reference [code table](#code-table) below

### With BBcode Format

```go
import "github.com/iskaa02/qalam/bbcode"
func main(){
	// this will output "a red text" with red in terminal
	// you can use bbcode.Sprintf() if you don't want to print yet
	bbcode.Printf("[red]a red %s[/red]","text")
}
```

<p align="right">(<a href="#top">back to top</a>)</p>

## BBcode Examples

### Multiple styles

```go
bbcode.Printf("[b s]bold and strikethrough[/b s]")
```

outputs:
~~**bold and strikethrough**~~

<br/>

### Composing styles

```go
bbcode.Printf("[b]bold only [i]bold and italic[/i][/b]")
```

outputs:
**bold only** **_bold and italic_**

<br/>

## Code Table

| Code            | Style                                                    |
| --------------- | -------------------------------------------------------- |
| b/bold          | **bold text**                                            |
| i/italic        | _italic text_                                            |
| u/underline     | underline text                                           |
| s/strikethrough | ~~strikethrough text~~                                   |
| #hexcode        | change the color                                         |
| bg#hexcode      | change the background color                              |
| red             | change the color to red                                  |
| yellow          | change the color to yellow                               |
| magenta         | change the color to magenta                              |
| white           | change the color to white                                |
| black           | change the color to black                                |
| green           | change the color to green                                |
| cyan            | change the color to cyan                                 |
| bg.(color)      | change the background color to any color mentioned above |

<p align="right">(<a href="#top">back to top</a>)</p>

## Gradient

### Import

```go
import "github.com/iskaa02/qalam/gradient"
```

### Usage

```go
g,err:=gradient.NewGradient("cyan","red")
// if the passed colors are not valid an error is returned
// supports all css valid colors
if err != nil{
}
g.Print("Example gradient from Cyan to Red")
```

outputs the following

![gradient Example](./doc/images/gradient_example1.png)

For more examples [checkout gradient Readme](https://github.com/iskaa02/qalam/blob/main/gradient/README.MD)

<p align="right">(<a href="#top">back to top</a>)</p>

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

## Notes

## Windows Support

If you're on Windows, use [Windows Terminal](https://github.com/microsoft/terminal) instead of cmd.exe.

#

## What does the Qalam mean

Qalam in Arabic means pencil

<br />

# Contact

Ismael Karim - ismael.karim000@gmail.com

Project Link: [https://github.com/iskaa02/qalam](https://github.com/iskaa02/qalam)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/iskaa02/qalam.svg?style=for-the-badge
[contributors-url]: https://github.com/iskaa02/qalam/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/iskaa02/qalam.svg?style=for-the-badge
[forks-url]: https://github.com/iskaa02/qalam/network/members
[stars-shield]: https://img.shields.io/github/stars/iskaa02/qalam.svg?style=for-the-badge
[stars-url]: https://github.com/iskaa02/qalam/stargazers
[issues-shield]: https://img.shields.io/github/issues/iskaa02/qalam.svg?style=for-the-badge
[issues-url]: https://github.com/iskaa02/qalam/issues
[license-shield]: https://img.shields.io/github/license/iskaa02/qalam.svg?style=for-the-badge
[license-url]: https://github.com/iskaa02/qalam/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
