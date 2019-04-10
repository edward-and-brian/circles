import { Colors } from '../themes';
import moment from 'moment';

const getTestMessages = () => [
  {
    title: 'Today',
    data: [
      {
        user: {
          name: 'Edward Bachoura',
          avatar: Colors.red,
        },
        date: new Date('March 12, 2019 14:23'),
        messages: ['Ok just left'],
      },
      {
        user: {
          name: 'Edward Bachoura',
          avatar: Colors.red,
        },
        date: new Date('March 12, 2019 13:14'),
        messages: [
          'I’m hungry',
          'Let’s go to centre point',
          'give me a like an hour, I’ll let you know when I leave',
          'Lol fine',
        ],
      },
      {
        user: {
          name: 'Natalya Bachoura',
          avatar: Colors.green,
        },
        date: new Date('March 12, 2019 13:11'),
        messages: ['You choose lol I narrowed it down'],
      },
      {
        user: {
          name: 'Edward Bachoura',
          avatar: Colors.red,
        },
        date: new Date('March 12, 2019 13:09'),
        messages: ['Brea is a nice middle area', 'Down for either of those'],
      },
      {
        user: {
          name: 'Natalya Bachoura',
          avatar: Colors.green,
        },
        date: new Date('March 12, 2019 13:02'),
        messages: [
          'Or centre point cafe',
          'That’s like a good halway point?',
          'We can meet at reborn coffee',
          'Yeah I’m good with either',
        ],
      },
      {
        user: {
          name: 'Edward Bachoura',
          avatar: Colors.red,
        },
        date: new Date('March 12, 2019 12:40'),
        messages: [
          'Ok should I come out to you? or do you want to meet at a Starbucks somewhere in the middle?',
        ],
      },
      {
        user: {
          name: 'Natalya Bachoura',
          avatar: Colors.green,
        },
        date: new Date('March 12, 2019 11:57'),
        messages: ['Sounds good ' + String.fromCodePoint(0x1f44d)],
      },
      {
        user: {
          name: 'Edward Bachoura',
          avatar: Colors.red,
        },
        date: new Date('March 12, 2019 11:54'),
        messages: ['let me shower real quick and we’ll talk', 'I’m down'],
      },
      {
        user: {
          name: 'Natalya Bachoura',
          avatar: Colors.green,
        },
        date: new Date('March 12, 2019 11:52'),
        messages: [
          'I’m studying all day so like if you want to meet somewhere to study/do whatever you gotta do for senior project or whatever I’m down but lmk :)) otherwise ima lock myself in the student center and try to finish everything. I’m good either way',
          'Hi',
        ],
      },
    ],
  },
  {
    title: 'Yesterday',
    data: [
      {
        user: {
          name: 'Edward Bachoura',
          avatar: Colors.red,
        },
        date: new Date('March 12, 2019 19:01'),
        messages: ['Let’s talk tomorrow', 'Perfect'],
      },
      {
        user: {
          name: 'Natalya Bachoura',
          avatar: Colors.green,
        },
        date: new Date('March 12, 2019 18:42'),
        messages: ['Yeah from 2 - 4'],
      },
      {
        user: {
          name: 'Edward Bachoura',
          avatar: Colors.red,
        },
        date: new Date('March 11, 2019 18:10'),
        messages: ['Are you free to study together tomorrow?'],
      },
    ],
  },
];

const getChats = () => [
  {
    name: 'Claudia Bachoura',
    avatar: Colors.c_blue,
    circles: [
      {
        name: 'Groceries',
        avatar: Colors.sunset1,
        lastMessage: {
          content: 'Brussel sprouts plzz',
          createdAt: moment().subtract(1, 'hours'),
        }
      },
      {
        name: 'Travel',
        avatar: Colors.sunset2,
        lastMessage: {
          content: 'Make sure that you are free February 21!',
          createdAt: moment().subtract(2, 'days'),
        }
      },
      {
        name: 'General',
        avatar: Colors.sunset3,
        lastMessage: {
          content: 'Have you picked up your gown yet?',
          createdAt: moment().subtract(5, 'days').subtract(9, 'hours'),
        }
      },
    ],
  },
  // {
  //   name: 'Brian Joerger',
  //   recentCircle: 'Circles',
  //   date: moment().subtract(2.2, 'hours'),
  //   circles: [1, 2, 3, 4, 5, 6, 7, 8],
  // },
  // {
  //   name: 'The Wild Ones',
  //   recentCircle: 'summer oregeon trip',
  //   date: moment().subtract(9, 'hours'),
  //   circles: [1, 2, 3],
  // },
  // {
  //   name: 'Natalya Bachoura',
  //   recentCircle: 'General',
  //   date: moment().subtract(18, 'hours'),
  //   circles: [1, 2, 3, 4, 5, 6, 7, 8],
  // },
  // {
  //   name: 'Jimmy Neutron',
  //   recentCircle: 'Plasma Weapons',
  //   date: moment().subtract(8, 'days'),
  //   circles: [1, 2, 3, 4, 5, 6, 7, 8],
  // },
];

export default {
  getTestMessages,
  getChats,
};
