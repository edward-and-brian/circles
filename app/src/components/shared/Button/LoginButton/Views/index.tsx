import React, { PureComponent } from 'react';
import { TouchableOpacity, View, Text } from 'react-native';
import styles from './styles';

export interface Props {
  text: string;
  color: string;
  onPress(): void;
}

class LoginButtonView extends PureComponent<Props> {
  render() {
    return (
      <TouchableOpacity
        style={[styles.container, { backgroundColor: this.props.color }]}
        onPress={this.props.onPress}
      >
        <Text style={[styles.buttonText, { color: this.props.fontColor }]}>
          {this.props.text}
        </Text>
      </TouchableOpacity>
    );
  }
}

export default LoginButtonView;
