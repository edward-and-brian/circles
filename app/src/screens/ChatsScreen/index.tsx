import React, { PureComponent } from 'react';
import { NavigationScreenProp, withNavigation } from 'react-navigation';
import ChatsScreenView from './Views';

export interface Props {
  navigation: NavigationScreenProp<any, any>;
}

class ChatsScreen extends PureComponent<Props> {
  onPressChat = (): void => {
    this.props.navigation.navigate('circle');
  };

  render() {
    return <ChatsScreenView onPressChat={this.onPressChat} />;
  }
}

export default withNavigation(ChatsScreen);
