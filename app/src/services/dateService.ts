class DateService {
  isInThisWeek(date) {
    const now = new Date();
    const firstDayOfWeek = new Date(now.setDate(now.getDate() - now.getDay()));
    firstDayOfWeek.setHours(0, 0, 0, 0);

    const lastDayOfWeek = new Date(firstDayOfWeek);
    lastDayOfWeek.setDate(firstDayOfWeek.getDate() + 6);
    lastDayOfWeek.setHours(23, 59, 59, 999);

    return date >= firstDayOfWeek && date <= lastDayOfWeek;
  }

  isInThisMonth(date) {
    const now = new Date();
    const firstDayOfMonth = new Date(now.getFullYear(), now.getMonth(), 1);
    const lastDayOfMonth = new Date(
      now.getFullYear(),
      now.getMonth() + 1,
      0,
      23,
      59,
      59,
      999
    );

    return date >= firstDayOfMonth && date <= lastDayOfMonth;
  }

  isInThisQuarter(date) {
    const now = new Date();
    const quarterStartMonth = Math.floor(now.getMonth() / 3) * 3;
    const firstDayOfQuarter = new Date(now.getFullYear(), quarterStartMonth, 1);
    const lastDayOfQuarter = new Date(
      now.getFullYear(),
      quarterStartMonth + 3,
      0,
      23,
      59,
      59,
      999
    );

    return date >= firstDayOfQuarter && date <= lastDayOfQuarter;
  }

  isInThisHalfYear(date) {
    const now = new Date();
    const year = now.getFullYear();

    let firstDayOfHalfYear, lastDayOfHalfYear;
    if (now.getMonth() < 6) {
      // First half of the year: January 1 to June 30
      firstDayOfHalfYear = new Date(year, 0, 1); // January 1
      lastDayOfHalfYear = new Date(year, 5, 30, 23, 59, 59, 999); // June 30
    } else {
      // Second half of the year: July 1 to December 31
      firstDayOfHalfYear = new Date(year, 6, 1); // July 1
      lastDayOfHalfYear = new Date(year, 11, 31, 23, 59, 59, 999); // December 31
    }

    return date >= firstDayOfHalfYear && date <= lastDayOfHalfYear;
  }

  isInThisYear(date) {
    const now = new Date();
    const firstDayOfYear = new Date(now.getFullYear(), 0, 1); // January 1
    const lastDayOfYear = new Date(now.getFullYear(), 11, 31, 23, 59, 59, 999); // December 31

    return date >= firstDayOfYear && date <= lastDayOfYear;
  }
}

const dateService = new DateService();
export default dateService;
