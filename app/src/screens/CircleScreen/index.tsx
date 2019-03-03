import React, { PureComponent } from 'react';
import { NavigationScreenProp, withNavigation } from 'react-navigation';
import CircleScreenView from './Views';

export interface Props {
  navigation: NavigationScreenProp<any, any>;
}

class CircleScreen extends PureComponent<Props> {
  onPressBack = (): void => {
    this.props.navigation.goBack();
  };

  render() {
    return <CircleScreenView onPressBack={this.onPressBack} />;
  }
}

export default withNavigation(CircleScreen);
