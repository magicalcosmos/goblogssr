import jstz from 'jstz';

const Time = {};
var now = new Date(); //当前日期
var nowDayOfWeek = now.getDay(); //今天本周的第几天
var nowDay = now.getDate(); //当前日
var nowMonth = now.getMonth(); //当前月
var nowYear = now.getYear(); //当前年
nowYear += (nowYear < 2000) ? 1900 : 0; //
var lastMonthDate = new Date(); //上月日期
lastMonthDate.setDate(1);
lastMonthDate.setMonth(lastMonthDate.getMonth() - 1);
var lastYear = lastMonthDate.getYear();
var lastMonth = lastMonthDate.getMonth();

/**
 * 格式化日期：yyyy-MM-dd
 * @param {*} date 
 */
export function formatDate(date) {
  let year = date.getUTCFullYear();
  let month = date.getUTCMonth() + 1;
  let day = date.getUTCDate();
  let hour = date.getUTCHours();
  let minute = date.getUTCMinutes();
  let second = date.getUTCSeconds();

  return `${year}-${compareWithZero(month)}-${compareWithZero(day)} ${compareWithZero(hour)}:${compareWithZero(minute)}:${compareWithZero(second)}`;
}

/**
 * 格式化日期：yyyy-MM-dd
 * @param {*} date 
 */
export function formatToGreenDate(date) {
  let year = date.getUTCFullYear();
  let month = date.getUTCMonth() + 1;
  let day = date.getUTCDate();
  let hour = date.getUTCHours();
  let minute = date.getUTCMinutes();
  let second = date.getUTCSeconds();

  return `${year}-${compareWithZero(month)}-${compareWithZero(day)}T${compareWithZero(hour)}:${compareWithZero(minute)}:${compareWithZero(second)}Z`;
}

export function compareWithZero(n) {
  return n > 9 ? n : `0${n}`;
}

/**
 * 获得某月的天数
 * @param {*} myMonth 
 */
export function getMonthDays(myMonth) {
  var monthStartDate = new Date(nowYear, myMonth, 1);
  var monthEndDate = new Date(nowYear, myMonth + 1, 1);
  var days = (monthEndDate - monthStartDate) / (1000 * 60 * 60 * 24);
  return days;
}
/**
 * 获得本季度的开始月份
 */
export function getQuarterStartMonth() {
  var quarterStartMonth = 0;
  if (nowMonth < 3) {
    quarterStartMonth = 0;
  }
  if (nowMonth > 2 && nowMonth < 6) {
    quarterStartMonth = 3;
  }
  if (nowMonth > 5 && nowMonth < 9) {
    quarterStartMonth = 6;
  }
  if (nowMonth > 8) {
    quarterStartMonth = 9;
  }
  return quarterStartMonth;
}
/**
 * 获得本周的开始日期
 */
Time.getWeekStartDate = function getWeekStartDate() {
  var weekStartDate = new Date(nowYear, nowMonth, nowDay - nowDayOfWeek);
  return formatDate(weekStartDate);
};
/**
 * 获得本周的结束日期
 */
Time.getWeekEndDate = function getWeekEndDate() {
  var weekEndDate = new Date(nowYear, nowMonth, nowDay + (6 - nowDayOfWeek));
  return formatDate(weekEndDate);
};
/**
 * 获得上周的开始日期
 */
Time.getLastWeekStartDate = function getLastWeekStartDate() {
  var weekStartDate = new Date(nowYear, nowMonth, nowDay - nowDayOfWeek - 7);
  return formatDate(weekStartDate);
};
/**
 * 获得上周的结束日期e
 */
Time.getLastWeekEndDate = function getLastWeekEndDate() {
  var weekEndDate = new Date(nowYear, nowMonth, nowDay - nowDayOfWeek - 1);
  return formatDate(weekEndDate);
};
/**
 * 获得本月的开始日期
 */
export function getMonthStartDate() {
  var monthStartDate = new Date(nowYear, nowMonth, 1);
  return formatDate(monthStartDate);
}
/**
 * 获得本月的结束日期
 */
export function getMonthEndDate() {
  var monthEndDate = new Date(nowYear, nowMonth, getMonthDays(nowMonth));
  return formatDate(monthEndDate);
}
/**
 * 获得上月开始时间
 */
export function getLastMonthStartDate() {
  var lastMonthStartDate = new Date(nowYear, lastMonth, 1);
  return formatDate(lastMonthStartDate);
}
/**
 * 获得上月结束时间
 */
export function getLastMonthEndDate() {
  var lastMonthEndDate = new Date(nowYear, lastMonth, getMonthDays(lastMonth));
  return formatDate(lastMonthEndDate);
}
/**
 * 获得本季度的开始日期
 */
export function getQuarterStartDate() {
  var quarterStartDate = new Date(nowYear, getQuarterStartMonth(), 1);
  return formatDate(quarterStartDate);
}
/**
 * 获得本季度的结束日期
 */
export function getQuarterEndDate() {
  var quarterEndMonth = getQuarterStartMonth() + 2;
  var quarterStartDate = new Date(nowYear, quarterEndMonth,
    getMonthDays(quarterEndMonth));
  return formatDate(quarterStartDate);
}
/**
 * 获取时间区域
 */
export function getTimezoneName() {
  const timezone = jstz.determine();
  if (timezone.name() === 'Asia/Tokyo') {
    return 'ja-JP';
  } else if (timezone.name() === 'Asia/Shanghai') {
    return 'zh-CN';
  }
  return 'zh-CN';
}

export default Time;
