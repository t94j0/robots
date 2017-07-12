# robots

# How to use
* Import the `robots` module: `import github.com/t94j0/robots`
* Call the `parseRobots` function: 

```
...
robotsString := "User-Agent: *\nDisallow: /private\nDisallow: /disallow"

robots := robots.ParseRobots(robotsString)

fmt.Println(robots.Disallow)
fmt.Println(robots.UserAgents)
...
```
