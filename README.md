Design Airline Check-in System
===

We are trying to implement the problem statement here. Thanks Arpit :)  <br/>
https://github.com/relogX/system-design-questions/blob/master/airline-checkin.md





# Problem Statement

When you book your tickets with an airline you are required to complete the payment and confirm your reservation. Once the reservation is complete then you can either optionally do a web check-in and confirm your seats or just before your departure do a physical check-in at the airport.

In this problem statement, let's design this web-check in system, where the passenger logs in to the system with the PNR, performs the seat selection and the gets the boarding pass. If the passenger tries to book a seat, already booked and assigned to the other passenger show an error message requesting passenger to re-select the seats.

![Relog Airline Check-in System](https://user-images.githubusercontent.com/4745789/138721841-3fc02879-7075-491a-9dcf-74011dba11e6.png)

# Requirements

<!--rs-->
*The problem statement is something to start with, be creative and dive into the product details and add constraints and features you think would be important.*
<!--re-->

## Core Requirements

- **one seat** can be assigned to only one passenger and once assigned the seat cannot be transferred
- assume all **100 people** boarding the plane are trying to make a selection of their seat at the same time
- the check-in should be as **fast** as possible
- when one passenger is booking a seat it should **not** lead to other passengers waiting

##  High Level Requirements
<!--hs-->
- make your high-level components operate with **high availability**
- ensure that the data in your system is **durable**, not matter what happens
- define how your system would behave while **scaling-up** and **scaling-down**
- make your system **cost-effective** and provide a justification for the same
- describe how **capacity planning** helped you made a good design decision
- think about how other services will interact with your service
<!--he-->

##  Micro Requirements
<!--ms-->
- ensure the data in your system is **never** going in an inconsistent state
- ensure your system is **free of deadlocks** (if applicable)
- ensure that the throughput of your system is not affected by **locking**, if it does, state how it would affect
<!--me-->

# Output

## Design Document
<!--ds-->
Create a **design document** of this system/feature stating all critical design decisions, tradeoffs, components, services, and communications. Also specify how your system handles at scale, and what will eventually become a chokepoint.

Do **not** create unnecessary components, just to make design look complicated. A good design is **always simple and elegant**. A good way to think about it is if you were to create a spearate process/machine/infra for each component and you will have to code it yourself, would you still do it?
<!--de-->

## Prototype

To understand the nuances and internals of this system, build a prototype that

- design a database schema for the airline check-in system
- build a simple interface allowing passenger to
    - view available seats
    - view unavailable seats
    - select a seat of their liking
    - upon successful booking, print their boarding pass
- simulate multiple passengers trying to book the same seats and handle the concurrency

