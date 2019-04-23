import React, { PureComponent } from 'react';
import NavigationService from '../../navigation/NavigationService';
import CircleScreenView from './Views';

export interface Props {}
interface State {}

class CircleScreen extends PureComponent<Props, State> {
  onPressBack() {
    NavigationService.back();
  }

  render() {
    return <CircleScreenView onPressBack={this.onPressBack} />;
  }
}

export default CircleScreen;
