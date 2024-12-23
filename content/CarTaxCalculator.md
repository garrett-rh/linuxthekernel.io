---
id: "CarTaxCalculator"
title: "Car Tax Calculator"
date: "2024-12-23"
summary: "The vehicle personal property tax calculator."
tags: ["Car", "Calculator", "Virginia", "Personal", "Property", "Tax"]
---

# Car Tax Calculator

I recently added a [tax calculator for vehicle personal property taxes in Virginia](/car_tax).
I'm big on budgeting and recently purchased a new vehicle. After calculating my tax bill for the third or fourth time I figured I'd make something to do it for me.
Currently, it only has Arlington County's data in it but feel free to [raise an issue on my Github](https://github.com/garrett-rh/linuxthekernel.io/issues) for me to add new localities.
This calculator doesn't store any information on what user data is entered. If you don't believe me, feel free to [check the source code](https://github.com/garrett-rh/linuxthekernel.io/)!

## Full Disclosure

I was too lazy to implement the financial math utilizing integers instead of floats so you might end up seeing some minor rounding errors.
Please only use this for rough estimates. At its worse, I was only seeing errors of a few pennies.
Also, this calculator is only for personal usage vehicles without taking into account any other tax modifiers like clean fuel vehicles.
