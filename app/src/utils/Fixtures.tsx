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
          createdAt: moment().subtract(22, 'minutes'),
        },
      },
      {
        name: 'Travel',
        avatar: Colors.sunset2,
        lastMessage: {
          content: 'Make sure that you are free February 21!',
          createdAt: moment()
            .subtract(2, 'days')
            .subtract(6, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset3,
        lastMessage: {
          content: 'Have you picked up your gown yet?',
          createdAt: moment()
            .subtract(5, 'days')
            .subtract(9, 'hours')
            .subtract(26, 'minutes'),
        },
      },
    ],
  },
  {
    name: 'Brian Joerger',
    avatar: Colors.c_green,
    circles: [
      {
        name: '402',
        avatar: Colors.sunset1,
        lastMessage: {
          content: 'Ready to present?',
          createdAt: moment()
            .subtract(2, 'hours')
            .subtract(22, 'minutes'),
        },
      },
      {
        name: 'da-jub-hoont',
        avatar: Colors.sunset2,
        lastMessage: {
          content: 'Battelfy interview went really well :o',
          createdAt: moment()
            .subtract(1, 'days')
            .subtract(18, 'hours')
            .subtract(13, 'minutes'),
        },
      },
      {
        name: 'Gamer',
        avatar: Colors.sunset3,
        lastMessage: {
          content: 'Gamer? Gamer and gamer? Gamer.',
          createdAt: moment()
            .subtract(2, 'days')
            .subtract(18, 'hours')
            .subtract(10, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset4,
        lastMessage: {
          content: "We don't talk anymore :' (",
          createdAt: moment()
            .subtract(10, 'days')
            .subtract(12, 'hours')
            .subtract(45, 'minutes'),
        },
      },
    ],
  },
  {
    name: 'Natalya Bachoura',
    avatar: Colors.c_red,
    circles: [
      {
        name: 'Drama',
        avatar: Colors.sunset1,
        lastMessage: {
          content: 'God, Jessica is such a fake friend, just tell your ',
          createdAt: moment()
            .subtract(2, 'days')
            .subtract(14, 'hours')
            .subtract(48, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset2,
        lastMessage: {
          content: "Yeah, I'm on my way home right now",
          createdAt: moment()
            .subtract(5, 'days')
            .subtract(4, 'hours')
            .subtract(2, 'minutes'),
        },
      },
    ],
  },
  {
    name: 'The Wild Ones',
    avatar: Colors.c_orange,
    circles: [
      {
        name: 'summer oregeon trip',
        avatar: Colors.sunset1,
        lastMessage: {
          content: 'Just finished making hotel reservations.',
          createdAt: moment()
            .subtract(4, 'days')
            .subtract(1, 'hours')
            .subtract(2, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset2,
        lastMessage: {
          content:
            'BAR CRAWL TIME. PICKING UP ALL YOU MOFOS RIGHT NOW, BETTER BE READY.',
          createdAt: moment()
            .subtract(12, 'days')
            .subtract(4, 'hours')
            .subtract(34, 'minutes'),
        },
      },
    ],
  },
  {
    name: 'Jimmy Neutron',
    avatar: Colors.c_yellow,
    circles: [
      {
        name: 'Plasma Weapons',
        avatar: Colors.sunset1,
        lastMessage: {
          content:
            "CHECK OUT MY LATEST INVENTION. Greatest. Mindblast. EVER. You'll never have to deal with those pesky birds again.",
          createdAt: moment()
            .subtract(16, 'days')
            .subtract(6, 'hours')
            .subtract(17, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset2,
        lastMessage: {
          content: "Let's be friends.",
          createdAt: moment()
            .subtract(62, 'days')
            .subtract(17, 'hours')
            .subtract(55, 'minutes'),
        },
      },
    ],
  },
  
  // DUPLICATES

  {
    name: 'Haley Flesh',
    avatar: Colors.c_blue,
    circles: [
      {
        name: 'Groceries',
        avatar: Colors.sunset1,
        lastMessage: {
          content: 'Brussel sprouts plzz',
          createdAt: moment().subtract(22, 'minutes'),
        },
      },
      {
        name: 'Travel',
        avatar: Colors.sunset2,
        lastMessage: {
          content: 'Make sure that you are free February 21!',
          createdAt: moment()
            .subtract(2, 'days')
            .subtract(6, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset3,
        lastMessage: {
          content: 'Have you picked up your gown yet?',
          createdAt: moment()
            .subtract(5, 'days')
            .subtract(9, 'hours')
            .subtract(26, 'minutes'),
        },
      },
    ],
  },
  {
    name: 'Seniors Citizes',
    avatar: Colors.c_green,
    circles: [
      {
        name: '402',
        avatar: Colors.sunset1,
        lastMessage: {
          content: 'Ready to present?',
          createdAt: moment()
            .subtract(2, 'hours')
            .subtract(22, 'minutes'),
        },
      },
      {
        name: 'da-jub-hoont',
        avatar: Colors.sunset2,
        lastMessage: {
          content: 'Battelfy interview went really well :o',
          createdAt: moment()
            .subtract(1, 'days')
            .subtract(18, 'hours')
            .subtract(13, 'minutes'),
        },
      },
      {
        name: 'Gamer',
        avatar: Colors.sunset3,
        lastMessage: {
          content: 'Gamer? Gamer and gamer? Gamer.',
          createdAt: moment()
            .subtract(2, 'days')
            .subtract(18, 'hours')
            .subtract(10, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset4,
        lastMessage: {
          content: "We don't talk anymore :' (",
          createdAt: moment()
            .subtract(10, 'days')
            .subtract(12, 'hours')
            .subtract(45, 'minutes'),
        },
      },
    ],
  },
  {
    name: 'Bernard Bachoura',
    avatar: Colors.c_red,
    circles: [
      {
        name: 'Drama',
        avatar: Colors.sunset1,
        lastMessage: {
          content: 'God, Jessica is such a fake friend, just tell your ',
          createdAt: moment()
            .subtract(2, 'days')
            .subtract(14, 'hours')
            .subtract(48, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset2,
        lastMessage: {
          content: "Yeah, I'm on my way home right now",
          createdAt: moment()
            .subtract(5, 'days')
            .subtract(4, 'hours')
            .subtract(2, 'minutes'),
        },
      },
    ],
  },
  {
    name: 'Catch these pants',
    avatar: Colors.c_orange,
    circles: [
      {
        name: 'summer oregeon trip',
        avatar: Colors.sunset1,
        lastMessage: {
          content: 'Just finished making hotel reservations.',
          createdAt: moment()
            .subtract(4, 'days')
            .subtract(1, 'hours')
            .subtract(2, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset2,
        lastMessage: {
          content:
            'BAR CRAWL TIME. PICKING UP ALL YOU MOFOS RIGHT NOW, BETTER BE READY.',
          createdAt: moment()
            .subtract(12, 'days')
            .subtract(4, 'hours')
            .subtract(34, 'minutes'),
        },
      },
    ],
  },
  {
    name: 'Timmy Turner',
    avatar: Colors.c_yellow,
    circles: [
      {
        name: 'Plasma Weapons',
        avatar: Colors.sunset1,
        lastMessage: {
          content:
            "CHECK OUT MY LATEST INVENTION. Greatest. Mindblast. EVER. You'll never have to deal with those pesky birds again.",
          createdAt: moment()
            .subtract(16, 'days')
            .subtract(6, 'hours')
            .subtract(17, 'minutes'),
        },
      },
      {
        name: 'General',
        avatar: Colors.sunset2,
        lastMessage: {
          content: "Let's be friends.",
          createdAt: moment()
            .subtract(62, 'days')
            .subtract(17, 'hours')
            .subtract(55, 'minutes'),
        },
      },
    ],
  },
];

export default {
  getTestMessages,
  getChats,
};
