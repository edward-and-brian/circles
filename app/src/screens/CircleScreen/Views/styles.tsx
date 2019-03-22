import { StyleSheet } from 'react-native';
import { Colors, Fonts, Scaled } from '../../../themes/';

export default StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.white,
  },
  conversationContainer: {
    flex: 1,
  },
  headerContainer: {
    height: Scaled.navbarHeight,
  },
});
