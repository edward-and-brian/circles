import React, { PureComponent } from 'react';
import { View, Text } from 'react-native';
import CircleNavbar from '../../../components/Base/Navbars/CircleNavbar/';
import styles from './styles';

export interface Props {
  onPressBack(): void;
}

class CircleScreenView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <CircleNavbar onPressBack={this.props.onPressBack} />
        <View style={styles.conversationContainer}>
          <Text style={styles.light}>Circle CIRCLES circles</Text>
          <Text style={styles.book}>Circle CIRCLES circles</Text>
          <Text style={styles.medium}>Circle CIRCLES circles</Text>
          <Text style={styles.heavy}>Circle CIRCLES circles</Text>
          <Text style={styles.black}>Circle CIRCLES circles</Text>
        </View>
        <View style={styles.footerContainer} />
      </View>
    );
  }
}

export default CircleScreenView;
