import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../themes';

export default StyleSheet.create({
  container: {
    shadowOffset: { width: -0.5, height: -1 },
    shadowOpacity: 0.8,
    backgroundColor: Colors.white,
    alignItems: 'center',
    paddingTop: 10,
  },
  messageInput: {
    width: '90%',
    height: 30,
    paddingHorizontal: 10,
    borderColor: Colors.gray,
    borderWidth: 1,
    borderRadius: 8,
  },
});
