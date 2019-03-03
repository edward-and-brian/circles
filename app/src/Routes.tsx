import { createStackNavigator } from 'react-navigation';
import ChatsScreen from '../src/screens/ChatsScreen';
import CircleScreen from '../src/screens/CircleScreen';

export default createStackNavigator(
  {
    chats: ChatsScreen,
    circle: CircleScreen,
  },
  {
    initialRouteName: 'chats',
    headerMode: 'none',
  },
);
