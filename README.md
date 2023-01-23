# Description

In this task we created a more appealing, interactive and intuitive web interface for our ascii-art program

# Authors

Martin Vahe (mvahe)

Henri Suokas (hsuokas)

# Usage

To open the website first run our code `go run . `

Type `localhost:8080` into your web browser.

Choose a font and type your text in the text box, and press "Submit".
For clearing the textbox and output box, press "Clear".


# Implementation

GET method is used to generate a basic website after using `go run .` in the terminal and opening `localhost:8080` in the browser.

POST method is used after clicking the `submit` button to send text and banner type to GO server, which will use this data to generate ascii art and send it to the main page via html templates.