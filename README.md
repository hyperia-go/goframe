# goframe
Scientific computing for Golang

[![Build Status](https://travis-ci.org/celsomilne/goframe.svg?branch=master)](https://travis-ci.org/celsomilne/goframe)

## Goals of GoFrame
There are several packages which exist in languages like Python, MATLAB and R which give them distinct advantages as scientific computing languages. These languages are popular among scientists and engineers for their readability and the ease with which they are learned.

However, among the primary issues with these languages, are their speed. Python and R both suffer from the limitations imposed by interpreter locks, making concurrency a practical impossibility. This limits their usefulnes when applied to applications at scale.

MATLAB overcomes several of these issues but comes with a hefty price tag associated. Commercial licenses can go for as much as $10,000 per machine and there is a steep learning curve when using MATLAB.

Additionally, these interpreted languages suffer in performance in how they handle

GoLang was developed in 2012 by Google and has been growing quickly in popularity. Its efficient concurrency and static typing give it several advantages over Python and R without sacrificing readability. GoLang is highly scalable with several large companies already implementing it in production-level deployment. Docker, Atlassian, Adobe, Github, Heroku, IBM, Keybase, Uber and Google all use Go to no small extent in production and it is built and maintained by engineers at Google and a swelling open source community who develop repositories like this one.

GoFrame hopes to bring the convenience and readability of numerical computing in Python and R to GoLang, giving data scientists and analysts the power and reliability they need to develop scalable and efficient data-based applications.
