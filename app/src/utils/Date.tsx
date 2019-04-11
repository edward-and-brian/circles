import moment from 'moment';

const toDisplayString = (dateTime: moment.Moment) => {
  const date = moment().set('hour', 0).set('minute', 0);
  if (dateTime.isAfter(date.subtract(1, 'days'))) {
    return dateTime.format('h:mm a');
  } else if (dateTime.isAfter(date.subtract(2, 'days'))) {
    return 'Yesterday';
  } else if (dateTime.isAfter(date.subtract(7, 'days'))) {
    return dateTime.format('dddd');
  } else {
    return dateTime.format('M/D/YY');
  }
};

export default {
  toDisplayString,
};
