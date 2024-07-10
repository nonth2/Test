# find_busiest_intersections.py

def find_busiest_intersections(intersections):
    if not intersections:
        return []
    
    max_vehicles = max(intersections.values())
    busiest_intersections = [intersection for intersection, count in intersections.items() if count == max_vehicles]
    
    return busiest_intersections


# Unit tests
import unittest

class TestFindBusiestIntersections(unittest.TestCase):
    def test_single_busiest_intersection(self):
        data = {'A': 10, 'B': 20, 'C': 15}
        result = find_busiest_intersections(data)
        self.assertEqual(result, ['B'])

    def test_multiple_busiest_intersections(self):
        data = {'A': 20, 'B': 20, 'C': 15}
        result = find_busiest_intersections(data)
        self.assertCountEqual(result, ['A', 'B'])

    def test_all_intersections_same(self):
        data = {'A': 10, 'B': 10, 'C': 10}
        result = find_busiest_intersections(data)
        self.assertCountEqual(result, ['A', 'B', 'C'])

    def test_empty_dictionary(self):
        data = {}
        result = find_busiest_intersections(data)
        self.assertEqual(result, [])

    def test_single_intersection(self):
        data = {'A': 10}
        result = find_busiest_intersections(data)
        self.assertEqual(result, ['A'])

if __name__ == '__main__':
    unittest.main()

