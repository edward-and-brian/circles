import React, { PureComponent } from 'react';
import {
  TextInput,
  Animated,
  View,
  TouchableOpacity,
  Image,
} from 'react-native';
import { Images } from '../../../../themes';
import styles from './styles';

export interface Props {
  keyboardPadding: Animated.Value;
  message: string;
  renderSendArrow: boolean;
  onMessageChange(newMessage: string): void;
}

class FooterView extends PureComponent<Props> {
  renderMessageInput() {
    return (
      <View style={styles.inputContainer}>
        <TextInput
          style={styles.messageInput}
          value={this.props.message}
          onChangeText={this.props.onMessageChange}
          placeholder="Message"
          multiline
        />
      </View>
    );
  }

  renderSendArrow() {
    return (
      <TouchableOpacity style={styles.arrowContainer}>
        <Image
          source={Images.sendArrow}
          style={styles.sendArrow}
          resizeMode="contain"
        />
      </TouchableOpacity>
    );
  }
  render() {
    return (
      <Animated.View
        style={[
          styles.container,
          { paddingBottom: this.props.keyboardPadding },
        ]}
      >
        {this.renderMessageInput()}
        {this.props.renderSendArrow && this.renderSendArrow()}
      </Animated.View>
    );
  }
}

export default FooterView;
