import React, { PureComponent } from 'react';
import { View, Text, Animated } from 'react-native';
import CircleNavbar from '../../../components/shared/Navbars/CircleNavbar/';
import CircleFooter from '../../../components/CircleFooter';
import styles from './styles';

export interface Props {
  onPressBack(): void;
  footerHeight: Animated.Value;
}

class CircleScreenView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <CircleNavbar
          onPressBack={this.props.onPressBack}
          style={styles.headerContainer}
        />
        <View style={styles.conversationContainer}>
          <Text style={styles.light}>Circle CIRCLES circles</Text>
          <Text style={styles.book}>Circle CIRCLES circles</Text>
          <Text style={styles.medium}>Circle CIRCLES circles</Text>
          <Text style={styles.heavy}>Circle CIRCLES circles</Text>
          <Text style={styles.black}>Circle CIRCLES circles</Text>
        </View>
        <CircleFooter height={this.props.footerHeight} />
      </View>
    );
  }
}

export default CircleScreenView;
