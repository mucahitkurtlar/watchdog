[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



  <h3 align="center">watchdog</h3>

  <p align="center">
    Home automation and security system
    <br />
    <br />
    <a href="https://github.com/mucahitkurtlar/watchdog/issues">Report Bug</a>
    ·
    <a href="https://github.com/mucahitkurtlar/watchdog/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project


<b>watchdog</b> can control lights, notify you about motions when you are not at home.


### Built With

Watchdog built with these libraries and frameworks:
* [ESP8266WiFi](https://arduino-esp8266.readthedocs.io/en/latest/esp8266wifi/readme.html)
* [WiFi](https://www.arduino.cc/en/Reference/WiFi)
* [WiFiClientSecure](https://github.com/espressif/arduino-esp32/tree/master/libraries/WiFiClientSecure)
* [Telegram API](https://core.telegram.org/)
* [Gorilla Mux](https://github.com/gorilla/mux/)
* [Go Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api/)


<!-- GETTING STARTED -->
## Getting Started

You can use this project with following steps.

### Prerequisites

* 2 x NodeMCU
* Raspberry Pi
* PIR Motion Sensor
* IRFZ44N
* NPN BJ Transistor (I used BC547)
* LED Strip
* 2 x 10K resistor


### Installation

1. Open Telegram app and search for BotFather.
2. Click or type "/start" and follow the inctructions.
3. Note the "bot token".
4. Now search for IDBot and do the same things at step 2.
5. Note the "id"
6. Clone the repo
    ```sh
    git clone https://github.com/mucahitkurtlar/watchdog.git
    ```
7. Create "secrets.go" file in "server/src/secrets" folder.
8. Define following consts
    ```Go
    // Enter the bot token down below
    const BotToken = "XXXXXXXXXX:XXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    // Enter your chat id from step 5 down below
    const ChatID = XXXXXXXXX
    ```
9. Run ```go build```  in server/src/secrets
10. Run ```go install```  in server/src
11. You are ready to deploy and run your server executable
12. Make sure all libraries and boards installed from Arduino IDE borad and library manager.
13. Create "secrets.hpp" files in led/ and lamp/ folders
14. Define following consts in both files
    ```Cpp
    // Enter your AP SSID down below
    #define _SSID "Your_SSID_here"
    // Enter your AP password down below
    #define _PASS "Your_pass_here"
    ```
15. Upload your sketchs to the boards
11. Send "/start" message to your bot and follow the instructions :smile:




<!-- LICENSE -->
## License

Distributed under the GPL-3.0. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Mücahit KURTLAR - [@mucahit_kurtlar](https://www.instagram.com/mucahit_kurtlar/)

Project Link: [https://github.com/mucahitkurtlar/watchdog](https://github.com/mucahitkurtlar/watchdog)




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/mucahitkurtlar/watchdog.svg?style=for-the-badge
[contributors-url]: https://github.com/mucahitkurtlar/watchdog/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/mucahitkurtlar/watchdog.svg?style=for-the-badge
[forks-url]: https://github.com/mucahitkurtlar/watchdog/network/members
[stars-shield]: https://img.shields.io/github/stars/mucahitkurtlar/watchdog.svg?style=for-the-badge
[stars-url]: https://github.com/mucahitkurtlar/watchdog/stargazers
[issues-shield]: https://img.shields.io/github/issues/mucahitkurtlar/watchdog.svg?style=for-the-badge
[issues-url]: https://github.com/mucahitkurtlar/watchdog/issues
[license-shield]: https://img.shields.io/github/license/mucahitkurtlar/watchdog.svg?style=for-the-badge
[license-url]: https://github.com/mucahitkurtlar/watchdog/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/mucahitkurtlar/
