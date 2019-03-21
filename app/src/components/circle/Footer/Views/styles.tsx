import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../themes';

const textInputHeight = Scaled.screen.width * 0.075;

export default StyleSheet.create({
  container: {
    flexDirection: 'row',
    shadowOffset: { width: -0.5, height: -1 },
    shadowOpacity: 0.8,
    backgroundColor: Colors.white,
    paddingTop: Scaled.screen.width * 0.02,
  },
  inputContainer: {
    flex: 8,
    alignItems: 'center',
  },
  messageInput: {
    height: textInputHeight,
    width: '95%',
    paddingHorizontal: Scaled.screen.width * 0.023,
    fontFamily: Fonts.book,
    fontSize: Scaled.fontSize.h10,
    borderColor: Colors.gray,
    borderWidth: 1,
    borderRadius: 8,
  },
  arrowContainer: {
    justifyContent: 'center',
    alignItems: 'center',
    width: textInputHeight,
    height: textInputHeight,
    borderRadius: textInputHeight / 2,
    marginRight: '2.5%',
    backgroundColor: Colors.blue,
  },
  sendArrow: {
    height: '70%',
    width: '70%',
    tintColor: Colors.white,
  },
});
