import React, { PureComponent } from 'react';
import { View, Text } from 'react-native';
import { Scaled } from '../../../../themes';
import AvatarButton from '../../../shared/Button/AvatarButton';
import styles from './styles';

export interface Props {}

class NavbarView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.titleContainer}>
          <Text style={styles.title}>Chats</Text>
        </View>
        <View style={styles.avatarContainer}>
          <AvatarButton diameter={Scaled.screen.height * 0.066} />
        </View>
      </View>
    );
  }
}

export default NavbarView;
