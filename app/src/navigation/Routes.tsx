import { createStackNavigator } from 'react-navigation';
import ChatsScreen from '../screens/ChatsScreen';
import CircleScreen from '../screens/CircleScreen';
import SignInScreen from '../screens/SignInScreen';

export default createStackNavigator(
  {
    chats: ChatsScreen,
    circle: CircleScreen,
    user: SignInScreen
  },
  {
    initialRouteName: 'user',
    headerMode: 'none',
  },
);
