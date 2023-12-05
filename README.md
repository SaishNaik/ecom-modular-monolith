**Ecom Modular Monolith** (WIP)

**Inspiration**
Having currently built ecom_clean_architecture, we found that we have upgraded quite 
a lot from layered architecture. But still there are some challenges. 
The major challenges are 
1) context switching. 
Whenever we add a new feature, we have
to go through and understand the whole project. Example: Lets say `orders` dev 
teams want to add a new feature. It has to go through the whole project which includes other 
domains such as products, users etc. Also the dependencies/libraries will be for the
whole project which will indeed add a learning curve as well. So what we have done with clean is removed 
coupling between layers and introduced coupling within the layer
2) Agile team issues: In a fast paced agile team, where features need to be
released quickly to the market and involves multiple teams working (orders team,
payment team etc), this will cause problems during merge as they are likely to 
touch common code
3) No flexibility of dependencies
For some domains, we might not need a orm or might need a different one. It makes
it difficult to add a new orm dependency as it will be extra work to integrate it and 
to convince for a change when there is already one orm. (Mostly people will try to be
uniform in terms of orm)
4) Folder/Codebase structure and Over use of  Interfaces
Sometimes we might not actually want to complicate things. Lets say its a simple service
We might not want to complicate it with too many interfaces. But when we adapt a 
clean architecture, we have to follow certain principles. Also codebase segregation
into multiple folders can also lead to complications



**Solution**
Vertical Slice Architecture/ Modular Monolith Architecture

1)Vertical Slice Architecture as the name suggests is a vertical slice/column
which will go through all layers. The principle of vertical slice is simple:
Decrease coupling inside the vertical slice/column and increase slicing among slices
What this practically means is that we divide our code based on business
capabilities/features rather than technical layers as done in clean architecture

2)Modular Monolith Architecture means dividing the codebase into different domains
using DDD. For each domain, we will have a separate module. By doing this, 
multiple teams can work together within the bounded context without affecting each other. 
You can decide your own architecture per module. Faster development is one of
the many benefits of this.

3)Also we will use headless architecture as well. Headless architecture means 
separating frontend from backend and connecting them via api. This has benefits
such as omnichannel support, Independent release of UI and faster feature 
development for ui by taking the ui of the backend.

Drawbacks of Solution
1) Although Modular Monolith helps us in application scalabilty by using load balancer, 
the database still is singular and can only be scaled up to certain extent
2) We cannot deploy domains/services separately
3) We cannot scale services/domains separately. (Think about festival sales)

Solution to Drawbacks:
Microservices


**Design:**
FR:
1) List Products
2) Add Products to cart
3) Checkout Cart and create Order
4) User Register and Signin

NFR:

Divide into 4 modules: product, orders, carts and users





