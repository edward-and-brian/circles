## 5.0	Requirements Specification

#### 5.1  Introduction

Circles is a messaging application that has a fair balance between functionality and ease of use. The most central idea behind our app that separates it from other messaging apps is our lightweight and intuitive take on chats, both group and individual. Circles organizes chats into "circles" in a way that makes it feel as simple as standard messaging apps with the functionality seamlessly tagged on. With this approach, Circles will have a more welcoming and pleasing design that encourages users to experiment with functionality rather than shy away from it. There are multiple other pieces of functionality that we plan to add, such as customizable user profiles and flexible notification settings, but these ideas are still in the works.

------

#### 5.2  CSCI Component Breakdown

CSCI Circles is composed of the following CSCs:

------

5.2.1	Mobile Application CSC -- with various functionalities for the user, revolving around sending
		messages

5.2.1.1		Landing page CSU -- intitial page to collect information from user

5.2.1.1.1			login/sign-up buttons -- intoduction to the application, from where users can either
				login with an existing circles account or create an account

5.2.1.2		Login/sign-up flow CSU -- prompt the user to create an account with circles, which uses
			their cell phone number as a primary key

5.2.1.2.1			Input screen module -- users inputs cell phone number to receive text with
				verification code to log in with

5.2.1.3		Chats page CSU -- Displays all open and existing chats that the user has

5.2.1.3.1 		Chats module -- All chats are displayed (both group and individual conversations)

5.2.1.3.2 		New chat module -- button used to start a new conversation with either a group or an
				individual. This will be a new page.

5.2.1.3.3 		Search module -- Allows users to search through their chats and circles

5.2.1.4		New chat page CSU -- Allows the user to create a new chat

5.2.1.4.1 		Search users module -- search for users to chat with and add them to the chat that you're
				creating

5.2.1.4.1 		Chat avatar module -- Allow the user to select a photo to set as the avatar of that chat

5.2.1.4.2 		Chat name module -- If it's a groupchat, a name and an avatar for the chat would be
				required. In this module, you would provide the name of the chat

5.2.1.5		Circle page CSU -- The page that shows the contents of a specific circle

5.2.1.5.1			Messages module -- displays all of the messages within that circle, which you can read by
				scrolling through the ScrollView

5.2.1.5.2			Text input module -- Text box through which you can type your new message to send to
				that circle

5.2.1.5.3			Header module -- Displays the title of the current circle, as well as the group chat to which
				it belongs and the avatar of that chat as well.

5.2.1.6		Profile settings CSU -- Page for user to check and change their user information and settings

------

5.2.2	Server CSC -- Monolithic application hosted on Docker (Linux)

5.2.2.1		Postgres Database -- contains user, chat, circle, and message data

5.2.2.2		GraphQL API -- interfaces with database to allow mobile application to pull data and trigger 		
			server-side functions

5.2.2.3		User Authentication -- authenticates requests from users against data in database

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
