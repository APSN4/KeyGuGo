# KeyGuGo

Это приложение для хранения приватных данных. Здесь используется шифрование
AES-256-CBC с использованием локального хранилища. Ваши данные в полной сохранности.  
<p align="center">
  <img width="256" height="256" src="./build/appicon.png">
</p>

![Screenshot](./assets/Screenshot.png)  

Функционал:
* Добавление заметок
* Удаление заметок
* Скопировать полностью содержимое заметки
* Редактировать заметку  

Приложение доступно на Windows x64, Linux x64, MacOS arm64.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
