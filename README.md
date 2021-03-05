## WARNING: This repo is currently under active developing, dependencies and itself will have breaking change frequently. If you are interested in this idea, please leave me a message!
# Streamline
## Introduction
This project aim to address a production issue: Collaboration is hard. It is hard to understand other's code. It is hard to debug a super long function especially it is not written by you. More importantly, even among a group of skilled programmer, mass collaboration is still not possible.

In this project I introduced a novel approach(not really) building a service. To demonstrate in a trivial way, I used an analog of factory and streamline.

## Basic idea
Ideally, data and procedures that manipulate the data should be separated. Like the streamline in a factory: parts move along the conveyor belt, each time it arrives at a worker(in our case, procedure that manipulate this kind of data), the worker apply some change on the part.

Fortunately, Golang is very suitable for this task. By using `interface`, we can define what kind of information the data have, and define what information a procedure need. If the data satisfy the requirement, we can apply the procedure to the data and guaranteed a defined behavior.

## Concepts
### DataDomain
A `data-domain` belong to a `ConveyorBelt`, `ConveyorBelt` take this `data-domain` to all the `procedures` in the `streamline` that this `ConveyorBelt` is executing.

`Data-domain` contains all data and intermediate values needed to process the request. Thus, applying compatible `procedure` to this `data-domain` should behave like a mathematical operator that is closed under this domain.
### Factory
A `factory` is used to produce new `streamlines` and manage their meta-data.

### Procedure
A `Procedure` is a function, defined as `func(*ConveyorBelt) int`. The return value is status code, currently following http status code standard, but likely to change in the future.

### Streamline
A `streamline` contains multiple `procedures`. It defines a series of operation and their sequence.
### ConveyorBelt
`ConveyorBelt` execute a `streamline`, taking a `data domain` with it, and controlled by a `context`.

## How
First, you need to think and define `data-domains`. A `data-domain` should be sufficient to hold all the data needed across the entire sequence of computation. `Data-domain` should implement getters for all fields, purpose of this is to use `interface` to validate compatibility in compiling time.

Second, create `procedures` you want to apply to your data. Each `procedure` should associate with an `interface` that defines all fields required by this `procedure`. The `interface` contains multiple getters, if the corresponding `data-domain` implements those getters, the interface assertion will success.

Third, use the given `factory` to create `streamlines`, which is simply a sequence of `procedures`.

Lastly, create a `ConveyorBelt`, give it a `streamline`, a `data-domain` and a context to control execution of this `ConveyorBelt`.

## Future work
Apply this design pattern to an actual project and validate this idea. Any critic or comment are welcome!
