"""
Solutions for the Google technical interview
"""


class Employee:
    def __init__(self, employee_id: str) -> None:
        self.employee_id = employee_id
        self.manager: Employee | None = None
        self.employees: list[Employee] = []


class Company:
    def __init__(self) -> None:
        self.employees: dict[str, Employee] = {}

    def manager(self, employee: str, manager: str) -> None:
        """
        Indicates that employee is a direct report of manager.
        """
        e, m = self.employees.get(employee), self.employees.get(manager)

        if not e:
            e = Employee(employee_id=employee)
            self.employees[employee] = e

        if not m:
            m = Employee(employee_id=manager)
            self.employees[manager] = m

        e.manager = m
        m.employees.append(e)

    def peer(self, employee: str, other: str) -> None:
        """
        Indicates that employee and other are peers (same manager).
        """
        e, o = self.employees.get(employee), self.employees.get(other)

        if not e:
            e = Employee(employee_id=employee)
            self.employees[employee] = e

        if not o:
            o = Employee(employee_id=other)
            self.employees[other] = o

        if e.manager:
            o.manager = e.manager
            e.manager.employees.append(o)
        elif o.manager:
            e.manager = o.manager
            o.manager.employees.append(e)

    def reports_to(self, employee: str, other: str) -> bool:
        """
        Returns True if employee reports directly or indirectly to other,
        otherwise returns False.
        """
        e, o = self.employees.get(employee), self.employees.get(other)

        if not e or not o:
            return False

        while e:
            if e == o:
                return True
            e = e.manager
        return False
