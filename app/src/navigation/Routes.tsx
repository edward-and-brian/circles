import { createStackNavigator } from 'react-navigation';
import ChatsScreen from '../screens/ChatsScreen';
import CircleScreen from '../screens/CircleScreen';

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
