* jiro

Tool to use together with rofi as a quick link to jira.

Edit the jira URL in the source. Compile it, and then use it like this:

```shell
firefox $(./jiro $(rofi -dmenu -p ">>> " -l 1))
```

Awesome!
