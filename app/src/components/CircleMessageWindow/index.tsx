import React, { PureComponent } from 'react';
import CircleMessageWindowView from './Views';
import Fixtures from '../../utils/Fixtures';

const conversation = Fixtures.getTestMessages();

export interface Props {}
interface State {}

class CircleMessageWindow extends PureComponent<Props, State> {
  render() {
    return <CircleMessageWindowView conversation={conversation} />;
  }
}

export default CircleMessageWindow;
