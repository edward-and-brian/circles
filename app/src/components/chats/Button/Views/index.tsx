import React, { PureComponent } from 'react';
import { TouchableOpacity, Text } from 'react-native';
import { Chat } from '../../ScrollList';
import styles from './styles';

export interface Props {
  chat: Chat;
  onPressChat(): void;
}

class ButtonView extends PureComponent<Props> {
  render() {
    return (
      <TouchableOpacity
        style={styles.container}
        onPress={this.props.onPressChat}
      >
        <Text style={styles.name}>{this.props.chat.name}</Text>
      </TouchableOpacity>
    );
  }
}

export default ButtonView;
