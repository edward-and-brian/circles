import React, { PureComponent } from 'react';
import { ScrollView, View, Text, Animated } from 'react-native';
import CircleNavbar from '../../../components/shared/Navbars/CircleNavbar/';
import CircleFooter from '../../../components/CircleFooter';
import styles from './styles';

export interface Props {
  onPressBack(): void;
  footerHeight: Animated.Value;
  message: string;
  onMessageChange(newMessage: string): void;
}

class CircleScreenView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <CircleNavbar
          onPressBack={this.props.onPressBack}
          style={styles.headerContainer}
        />
        <ScrollView style={styles.conversationContainer}>
          <Text style={styles.light}>Circle CIRCLES circles</Text>
          <Text style={styles.book}>Circle CIRCLES circles</Text>
          <Text style={styles.medium}>Circle CIRCLES circles</Text>
          <Text style={styles.heavy}>Circle CIRCLES circles</Text>
          <Text style={styles.black}>Circle CIRCLES circles</Text>
        </ScrollView>
        <CircleFooter
          height={this.props.footerHeight}
          message={this.props.message}
          onMessageChange={this.props.onMessageChange}
        />
      </View>
    );
  }
}

export default CircleScreenView;
