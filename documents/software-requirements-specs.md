## 5.0	Requirements Specification

#### 5.1  Introduction

Circles is a messaging application that has a fair balance between functionality and ease of use. The most central idea behind our app that separates it from other messaging apps is our lightweight and intuitive take on chats, both group and individual. Circles organizes chats into "circles" in a way that makes it feel as simple as standard messaging apps with the functionality seamlessly tagged on. With this approach, Circles will have a more welcoming and pleasing design that encourages users to experiment with functionality rather than shy away from it. There are multiple other pieces of functionality that we plan to add, such as customizable user profiles and flexible notification settings, but these ideas are still in the works.

------

#### 5.2  CSCI Component Breakdown

CSCI Circles is composed of the following CSCs:

------

5.2.1	iOS Application CSC -- with various functionalities for both Buyers and Sellers on the app

5.2.1.1		Landing page CSU -- initial page loaded on app startup

5.2.1.1.1			Loading animation module -- elaborately designed loading animation

5.2.1.1.2			login/sign-up buttons -- allows users to begin the sign up or login process

5.2.1.2		Login/sign-up flow CSU -- general flow for users logging in to the app

5.2.1.2.1			Input screen module -- users input username, password, etc information

5.2.1.3		Browse/Main page CSU -- contains entire catalog of listings, along with other specific filters

5.2.1.3.1 		Featured module -- displays featured products with images

5.2.1.3.2 		Brands module -- displays all brands and allows you to filter by each

5.2.1.3.3 		Trending module -- displays all trending products that are listed

5.2.1.3.4 		Highest Offers module -- displays the highest offers currently made

5.2.1.3.5 		Categories module -- displays all Categories and allows you to filter by each

5.2.1.3.6 		Shop all module -- displays all products listed

5.2.1.4		Product page CSU -- contains all information and functionality for a given product

5.2.1.4.1 		Product info module -- displays prices, sizes, images, and etc info on the product

5.2.1.4.1 		Buy now module -- allows the user to purchase the given object

5.2.1.4.2 		Sell module -- allows sellers to list the given product for sale

5.2.1.5		Purchase/offer flow CSU -- show and retrieve necessary purchase info up until purchase 

5.2.1.5.1			Purchase info module -- displays costs, and other information to user

5.2.1.5.2			User input module -- takes in shipping address, and card info from user

5.2.1.5.3			Make offer module -- user creates new offer with the given information

5.2.1.5.4			Buy now module -- user completes purchase with the given information

5.2.1.6		Listing flow CSU -- seller creates a new listing for the product with the given info

5.2.1.6.1			Information input module -- takes in input for the seller's new listing

5.2.1.7		Notifications page CSU -- displays user's notifications

5.2.1.7.1			List all module -- lists all notifications for the given user, and whether they are new or not

5.2.1.8		Search/filter page CSU -- allows user to search for or filter through products

5.2.1.8.1			Search module -- takes in a string to search through products

5.2.1.8.2			Filter module -- takes in multiple filter options to filter through products

5.2.1.8.3			List desired output module -- displays products for the given search or filter

5.2.1.9		User page CSU -- allows user to search for or filter through products

5.2.1.9.1			Settings module -- allows user to view and edit their settings	

5.2.1.9.1			Profile module -- allows user to view and edit their profile information	

5.2.1.9.1			My orders module -- allows user to see all of their orders, and any info regarding them

5.2.1.9.2			My listings module -- allows user to see all of their listings, and any info regarding them

5.2.1.9.3			My offers module -- allows user to see their standing offers for products

5.2.1.9.4			My favorites module -- allows users to see and pick favorite products

------

5.2.2	Serverless backend CSC -- hosted on AWS for scalability, reliability, and variety of functionality

5.2.2.1		AWS built-in functionality CSU -- AWS's built in systems

5.2.2.1.1			AWS Cognito module -- AWS's built in user management system

5.2.2.1.2			AWS RDS module -- AWS's built in database hosting system

5.2.2.1.3			AWS Elasticache module -- AWS's built in cache hosting system

5.2.2.1.4			AWS S3 buckets module -- AWS's built in miscellaneous storage system

5.2.2.1.5			AWS SNS module -- AWS's built in micro-service communication system

5.2.2.1.6			AWS Lambda module -- AWS's built in backend interfacing system

5.2.2.2		Golang backend CSU -- provides backend requests to front-end through AWS Lambda

5.2.2.2.1			Inventory service module -- provides all functionality for our inventory system

5.2.2.2.2			Shipping service module -- provides all functionality for our shipping system

5.2.2.2.3			User service module -- provides all functionality for our user system

5.2.2.2.4			Analytics service module -- provides all functionality for our analytics system

------

#### 5.3  Functional Requirements by CSC

5.3.1    Mobile Application shall allow users to view their current chats, in order of recent activity.

5.3.2    Mobile Application shall allow users to view a chat's circles upon hovering the chat

5.3.3    Mobile Application shall provide to the user a view of a circle's messages upon request

5.3.4    Mobile Application shall allow users to edit notification settings for given chats and circles

5.3.5    Mobile Application shall send notifications to users dependent on their notification settings

5.3.6    Server shall provide an API with which the mobile appliation can retrieve various data

5.3.7    Server shall manage data in the database securely and effectively

5.3.8    Server shall properly authenticate users when processing a request

5.3.9    Server shall handle message sending and receiving in a secure and effective manner

------

#### 5.4  Performance Requirements by CSC

5.4.1    Mobile Application shall run seamlessly while retrieving various data

5.4.2    Mobile Application shall retrieve a user's initial list of chats within 1 second of startup

5.4.2    Mobile Application shall retrieve a chat's list of circles before the user can choose it's circles

5.4.3    Server shall retrieve the proper data that is requested

5.4.4    Server shall handle requests in a timely manner

5.4.5    Server shall gracefully handle unavoidable server errors