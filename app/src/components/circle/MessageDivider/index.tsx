import React, { PureComponent } from 'react';
import MessageDividerView from './Views';

export interface Props {
  label: string;
}
interface State {}

class MessageDivider extends PureComponent<Props, State> {
  render() {
    return <MessageDividerView {...this.props} />;
  }
}

export default MessageDivider;
