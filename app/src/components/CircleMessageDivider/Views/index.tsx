import React, { PureComponent } from 'react';
import { View, Text } from 'react-native';
import styles from './styles';

export interface Props {
  label: string;
}

class CircleMessageWindow extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.leftSpacer} />
        <Text style={styles.label}>{this.props.label}</Text>
        <View style={styles.rightSpacer} />
      </View>
    );
  }
}

export default CircleMessageWindow;
