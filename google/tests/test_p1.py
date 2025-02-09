from unittest import TestCase

from src.p1 import Company


class TestCompany(TestCase):
    def setUp(self):
        self.company = Company()

    def test_manager_basic(self):
        self.company.manager("alice", "bob")
        assert self.company.reports_to("alice", "bob")
        assert not self.company.reports_to("bob", "alice")

    def test_manager_chain(self):
        self.company.manager("alice", "bob")
        self.company.manager("bob", "charlie")
        assert self.company.reports_to("alice", "charlie")
        assert not self.company.reports_to("charlie", "alice")

    def test_peer_basic(self):
        self.company.manager("alice", "charlie")
        self.company.peer("alice", "bob")
        assert self.company.reports_to("alice", "charlie")
        assert self.company.reports_to("bob", "charlie")

    def test_complex_hierarchy(self):
        self.company.manager("dev_1", "tech_lead")
        self.company.manager("dev_2", "tech_lead")
        self.company.manager("tech_lead", "cto")
        self.company.manager("designer", "creative_director")
        self.company.manager("creative_director", "cto")

        assert self.company.reports_to("dev_1", "cto")
        assert self.company.reports_to("dev_2", "cto")
        assert self.company.reports_to("designer", "cto")
        assert not self.company.reports_to("cto", "dev_1")
        assert not self.company.reports_to("dev_1", "designer")

    def test_peer_existing_manager(self):
        self.company.manager("alice", "charlie")
        self.company.peer("alice", "bob")
        assert self.company.reports_to("bob", "charlie")
