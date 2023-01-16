# go-bot
**__Make sur you have <a href="https://go.dev/dl/">GoLang</a> installed before doing any of this__**

------------------------

The `main.go` script reads the config.json file on the root of the project, the structure of this file should be like this:

```json
{
    "Token":"YOUR BOT TOKEN"
}
```

This way the code reads the Token field and uses it to create the Discord session, this way you can keep the token safe and not hardcoded in the code.
<br><br>
**❓ Please keep in mind that this is just an example, and you can add more features to your bot and also you can add more fields to the `config.json` file to keep other configuration options. ❓**<br><br>
<img src="https://github.com/PwnHash/go-bot/blob/main/img/logo.png" alt="GoLang">
