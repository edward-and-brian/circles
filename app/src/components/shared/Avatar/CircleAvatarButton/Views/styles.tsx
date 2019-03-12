import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../../themes';

export default StyleSheet.create({
  container: {
    flex: 3,
    justifyContent: 'center',
    alignItems: 'center',
  },
  avatar: {
    height: Scaled.screen.height * 0.055,
    width: Scaled.screen.height * 0.055,
    borderRadius: 10000,
    backgroundColor: Colors.yellow,
  },
});
