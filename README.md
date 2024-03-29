# Alchemist

Alchemist is a non-deterministic programming language based on chemical reaction networks. [More information here](https://esolangs.org/wiki/Alchemist). This implementation is written in Go, the original implementation can be found [here](https://github.com/bforte/Alchemist). I would suggest reading [the original implementation's README](https://github.com/bforte/Alchemist/blob/master/README.md) to learn about how this particular language works.

This particular implementation has an API, so it can be executed from other Go code. The documentation can be found [here](https://godoc.org/github.com/deanveloper/alchemist).

## Feature List

 * [x] Working API
 * [x] CLI app to run Alchemy files
 * [ ] Write tests


## CLI Installation

Make sure you have [Go](https://golang.org/dl/) installed, and run the following commands:

```
go get github.com/deanveloper/alchemist/cmd
mv ~/go/bin/cmd ~/go/bin/alchemist
```

And make sure you have `~/go/bin` added to your `PATH`.

## CLI Usage

```
Usage: ./alchemist [flags]
Flags:
  -i string
        input file to read rules from
  -r string
        rules to execute with, overrides -i
  -u string
        universe of atoms to start with
```

<details><summary>Examples of rule sets:</summary>

### Hello world
<pre>
_->Out_"Hello, World!"
</pre>

### Hello world (using `!` to determine input universe)
<pre>
x->Out_"Hello, World!"!x
</pre>

### Countdown
<pre>
_->5x
x->Out_x
0x->Out_"Liftoff"
</pre>

### Adder
<pre>
_ -> a+reqX+Out_"Input 1:"+In_x+Out_"Input 2:"+In_y
x -> z
y -> z
a+0x+0y -> Out_z
</pre>

### Fibonacci
<pre>
_ -> Out_"Enter how many numbers you wanna see:"+In_loop+b+setNext+Out_""+Out_"Fibonacci:"+Out_a+Out_b

loop+a+setNext -> loop+next+setNext
loop+b+setNext -> loop+next+setNext+saveB
loop+0a+0b+setNext -> Out_next+setA

loop+setA+saveB -> loop+setA+a
loop+setA+0saveB -> loop+setB

loop+setB+next -> loop+setB+b
loop+setB+0next -> loop+setNext
</pre>
</details>
