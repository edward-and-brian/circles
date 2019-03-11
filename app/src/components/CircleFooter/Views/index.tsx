import React, { PureComponent } from 'react';
import { TextInput, Animated } from 'react-native';
import styles from './styles';

export interface Props {
  height: Animated.Value;
  message: string;
}

class CircleFooterView extends PureComponent<Props> {
  renderMessageInput() {
    return <TextInput style={styles.messageInput} placeholder="Message" />;
  }

  render() {
    return (
      <Animated.View style={[styles.container, { height: this.props.height }]}>
        {this.renderMessageInput()}
      </Animated.View>
    );
  }
}

export default CircleFooterView;
