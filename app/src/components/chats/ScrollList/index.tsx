import React, { PureComponent } from 'react';
import ScrollListView from './Views';
import moment from 'moment';

export interface Chat {
  name: string;
  recentCircle: string;
  date: moment.Moment;
  circles: number[];
}

export interface Props {
  chats: Chat[];
  onPressChat(): void;
}

class ScrollList extends PureComponent<Props> {
  render() {
    return <ScrollListView {...this.props} />;
  }
}

export default ScrollList;
