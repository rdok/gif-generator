# gif-generator

[![CI/CD](https://github.com/rdok/gif-generator/workflows/CI/CD/badge.svg)](https://github.com/rdok/gif-generator/actions?query=workflow%3ACI%2FCD)
[![testing-site](https://img.shields.io/badge/testing-grey?style=flat-square&logo=amazon-aws)](https://testing-gif-generator.rdok.co.uk/)
[![production-site](https://img.shields.io/badge/production-blue?style=flat-square&logo=amazon-aws)](https://gif-generator.rdok.co.uk/)

Customizable: background color, line color, canvas size

https://gif-generator.rdok.co.uk/?background-color=4272f5&line-color=000000&image-size=300

AWS Lambda process the GO binary, and AWS API Gateway exposes the request & response.

### Development
See `Makefile`