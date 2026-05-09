/**
 * 公历 → 农历转换
 * 基于经典 lookup table 算法（覆盖 1900-2100 年）
 */

// 农历数据：每个数字编码了一年中的农历信息
// bit 0-3: 闰月（0=无闰月）
// bit 4-15: 12 个月的大小月（0=29天, 1=30天）
// bit 16-19: 闰月大小（0=29天, 1=30天）
// bit 20+: 高四位未用
const LUNAR_INFO = [
  0x04bd8, 0x04ae0, 0x0a570, 0x054d5, 0x0d260, 0x0d950, 0x16554, 0x056a0, 0x09ad0, 0x055d2, // 1900-1909
  0x04ae0, 0x0a5b6, 0x0a4d0, 0x0d250, 0x1d255, 0x0b540, 0x0d6a0, 0x0ada2, 0x095b0, 0x14977, // 1910-1919
  0x04970, 0x0a4b0, 0x0b4b5, 0x06a50, 0x06d40, 0x1ab54, 0x02b60, 0x09570, 0x052f2, 0x04970, // 1920-1929
  0x06566, 0x0d4a0, 0x0ea50, 0x16a95, 0x05ad0, 0x02b60, 0x186e3, 0x092e0, 0x1c8d7, 0x0c950, // 1930-1939
  0x0d4a0, 0x1d8a6, 0x0b550, 0x056a0, 0x1a5b4, 0x025d0, 0x092d0, 0x0d2b2, 0x0a950, 0x0b557, // 1940-1949
  0x06ca0, 0x0b550, 0x15355, 0x04da0, 0x0a5b0, 0x14573, 0x052b0, 0x0a9a8, 0x0e950, 0x06aa0, // 1950-1959
  0x0aea6, 0x0ab50, 0x04b60, 0x0aae4, 0x0a570, 0x05260, 0x0f263, 0x0d950, 0x05b57, 0x056a0, // 1960-1969
  0x096d0, 0x04dd5, 0x04ad0, 0x0a4d0, 0x0d4d4, 0x0d250, 0x0d558, 0x0b540, 0x0b6a0, 0x195a6, // 1970-1979
  0x095b0, 0x049b0, 0x0a974, 0x0a4b0, 0x0b27a, 0x06a50, 0x06d40, 0x0af46, 0x0ab60, 0x09570, // 1980-1989
  0x04af5, 0x04970, 0x064b0, 0x074a3, 0x0ea50, 0x06b58, 0x05ac0, 0x0ab60, 0x096d5, 0x092e0, // 1990-1999
  0x0c960, 0x0d954, 0x0d4a0, 0x0da50, 0x07552, 0x056a0, 0x0abb7, 0x025d0, 0x092d0, 0x0cab5, // 2000-2009
  0x0a950, 0x0b4a0, 0x0baa4, 0x0ad50, 0x055d9, 0x04ba0, 0x0a5b0, 0x15176, 0x052b0, 0x0a930, // 2010-2019
  0x07954, 0x06aa0, 0x0ad50, 0x05b52, 0x04b60, 0x0a6e6, 0x0a4e0, 0x0d260, 0x0ea65, 0x0d530, // 2020-2029
  0x05aa0, 0x076a3, 0x096d0, 0x04afb, 0x04ad0, 0x0a4d0, 0x1d0b6, 0x0d250, 0x0d520, 0x0dd45, // 2030-2039
  0x0b5a0, 0x056d0, 0x055b2, 0x049b0, 0x0a577, 0x0a4b0, 0x0aa50, 0x1b255, 0x06d20, 0x0ada0, // 2040-2049
  0x14b63, 0x09370, 0x049f8, 0x04970, 0x064b0, 0x168a6, 0x0ea50, 0x06aa0, 0x1a6c4, 0x0aae0, // 2050-2059
  0x092e0, 0x0d2e3, 0x0c960, 0x0d557, 0x0d4a0, 0x0da50, 0x05d55, 0x056a0, 0x0a6d0, 0x055d4, // 2060-2069
  0x052d0, 0x0a9b8, 0x0a950, 0x0b4a0, 0x0b6a6, 0x0ad50, 0x055a0, 0x0aba4, 0x0a5b0, 0x052b0, // 2070-2079
  0x0b273, 0x06930, 0x07337, 0x06aa0, 0x0ad50, 0x14b55, 0x04b60, 0x0a570, 0x054e4, 0x0d160, // 2080-2089
  0x0e968, 0x0d520, 0x0daa0, 0x16aa6, 0x056d0, 0x04ae0, 0x0a9d4, 0x0a4d0, 0x0d150, 0x0f252, // 2090-2099
  0x0d520, // 2100
]

// 十二个月的中文名
const LUNAR_MONTH_NAMES = [
  '正月', '二月', '三月', '四月', '五月', '六月',
  '七月', '八月', '九月', '十月', '冬月', '腊月',
]

// 日期的中文名（1-30）
const LUNAR_DAY_NAMES = [
  '初一', '初二', '初三', '初四', '初五', '初六', '初七', '初八', '初九', '初十',
  '十一', '十二', '十三', '十四', '十五', '十六', '十七', '十八', '十九', '二十',
  '廿一', '廿二', '廿三', '廿四', '廿五', '廿六', '廿七', '廿八', '廿九', '三十',
]

// 星期中文名
const WEEK_NAMES = ['周日', '周一', '周二', '周三', '周四', '周五', '周六']

/**
 * 获取农历年的总天数
 */
function lunarYearDays(year: number): number {
  let sum = 348 // 12个月 × 29天
  const code = LUNAR_INFO[year - 1900]
  for (let i = 0x8000; i > 0x8; i >>= 1) {
    sum += (code & i) ? 1 : 0
  }
  return sum + leapMonthDays(year)
}

/**
 * 获取农历年闰月的天数
 */
function leapMonthDays(year: number): number {
  const code = LUNAR_INFO[year - 1900]
  return (code & 0xf000) ? ((code & 0x10000) ? 30 : 29) : 0
}

/**
 * 获取农历年闰月（0=无闰月）
 */
function leapMonth(year: number): number {
  return LUNAR_INFO[year - 1900] & 0xf
}

/**
 * 获取农历年指定月的天数
 */
function monthDays(year: number, month: number): number {
  return (LUNAR_INFO[year - 1900] & (0x10000 >> month)) ? 30 : 29
}

/**
 * 公历 → 农历转换
 */
function solarToLunar(year: number, month: number, day: number): {
  lunarYear: number
  lunarMonth: number
  lunarDay: number
  isLeap: boolean
} {
  // 基准日：1900年1月31日 = 农历庚子年正月初一
  const baseDate = new Date(1900, 0, 31)
  const targetDate = new Date(year, month - 1, day)
  let offset = Math.floor((targetDate.getTime() - baseDate.getTime()) / 86400000)

  // 从 1900 年开始遍历
  let lunarYear = 1900
  while (lunarYear < 2101 && offset > 0) {
    const days = lunarYearDays(lunarYear)
    offset -= days
    lunarYear++
  }
  if (offset < 0) {
    offset += lunarYearDays(lunarYear - 1)
    lunarYear--
  }

  // 现在 offset 是当年正月初一的偏移量（0=正月初一）
  const leap = leapMonth(lunarYear)
  let isLeap = false
  let lunarMonth = 1

  for (let i = 1; i <= 12; i++) {
    const days = monthDays(lunarYear, i)
    if (offset < days) {
      lunarMonth = i
      break
    }
    offset -= days

    // 如果有闰月
    if (leap > 0 && i === leap) {
      const leapDays = leapMonthDays(lunarYear)
      if (offset < leapDays) {
        lunarMonth = i
        isLeap = true
        break
      }
      offset -= leapDays
    }
    lunarMonth++
  }

  if (lunarMonth > 12) {
    lunarMonth = 12
  }

  return {
    lunarYear,
    lunarMonth,
    lunarDay: offset + 1,
    isLeap,
  }
}

export function useLunarCalendar() {
  function getLunarDateString(date: Date): string {
    const y = date.getFullYear()
    const m = date.getMonth() + 1
    const d = date.getDate()
    const w = date.getDay()

    const lunar = solarToLunar(y, m, d)
    const lunarMonthName = (lunar.isLeap ? '闰' : '') + LUNAR_MONTH_NAMES[lunar.lunarMonth - 1]
    const lunarDayName = LUNAR_DAY_NAMES[lunar.lunarDay - 1]

    // "05月09日 周六 三月廿三"
    const mm = String(m).padStart(2, '0')
    const dd = String(d).padStart(2, '0')
    return `${mm}月${dd}日 ${WEEK_NAMES[w]} ${lunarMonthName}${lunarDayName}`
  }

  return { getLunarDateString }
}
