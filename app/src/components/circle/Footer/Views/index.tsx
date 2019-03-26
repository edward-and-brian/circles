import React, { PureComponent } from 'react';
import {
  TextInput,
  Animated,
  View,
  TouchableOpacity,
  Image,
} from 'react-native';
import styles from './styles';
import { Images } from '../../../../themes';

export interface Props {
  height: Animated.Value;
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
          onChangeText={this.props.onMessageChange}
          value={this.props.message}
          placeholder="Message"
          enablesReturnKeyAutomatically
          returnKeyType="send"
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
      <Animated.View style={[styles.container, { height: this.props.height }]}>
        {this.renderMessageInput()}
        {this.renderSendArrow()}
        {/* {this.props.renderSendArrow && this.renderSendArrow()} */}
      </Animated.View>
    );
  }
}

export default FooterView;
