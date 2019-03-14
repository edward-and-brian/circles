import { Colors } from '../themes';

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

export default {
  getTestMessages,
};
