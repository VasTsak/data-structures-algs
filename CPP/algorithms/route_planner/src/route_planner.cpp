#include "route_planner.h"
#include <algorithm>
using std::sort;
using std::reverse;
RoutePlanner::RoutePlanner(RouteModel &model, float start_x, float start_y, float end_x, float end_y): m_Model(model) {
    // Convert inputs to percentage:
    start_x *= 0.01;
    start_y *= 0.01;
    end_x *= 0.01;
    end_y *= 0.01;

    // Use the m_Model.FindClosestNode method to find the closest nodes to the starting and ending coordinates.
    // Store the nodes you find in the RoutePlanner's start_node and end_node attributes.
	start_node = &m_Model.FindClosestNode(start_x, start_y);
    end_node = &m_Model.FindClosestNode(end_x, end_y);
}


// Implement the CalculateHValue method.

float RoutePlanner::CalculateHValue(RouteModel::Node const *node) {
	return node->distance( *(this->end_node) );
}

// Create AddNeighbors method to expand the current node by adding all unvisited neighbors to the open list
void RoutePlanner::AddNeighbors(RouteModel::Node *current_node) {
    // Use the FindNeighbors() method of the current_node to populate current_node.neighbors vector with all the neighbors.
    current_node->RouteModel::Node::FindNeighbors();
    // For each node in current_node.neighbors, set the parent, the h_value, the g_value. 
    for (RouteModel::Node* neighbor : current_node->neighbors) {
        neighbor->parent = current_node;
        neighbor->g_value = current_node->g_value + current_node->distance(*neighbor);
        // Use CalculateHValue below to implement the h-Value calculation.
        neighbor->h_value = this->CalculateHValue(neighbor);
        // For each node in current_node.neighbors, add the neighbor to open_list and set the node's visited attribute to true.
        this->open_list.emplace_back(neighbor); // There is no disadvantage to using emplace_back and using it can sometimes provide additional efficiency unlike push_back.
        neighbor->visited = true;
    }
}

// Create Compare method to sort the open_list according to the sum of the h value and g value.

bool RoutePlanner::Compare(const RouteModel::Node *node_a, const RouteModel::Node *node_b) {
  float f1 = node_a->g_value + node_a->h_value ;
  float f2 = node_b->g_value + node_b->h_value ;
  return f1 > f2; 
}

// Create NextNode method to sort the open list and return the next node.

RouteModel::Node *RoutePlanner::NextNode() {
    sort(this->open_list.begin(), this->open_list.end(), RoutePlanner::Compare);
    // Create a pointer to the node in the list with the lowest sum.
    RouteModel::Node* node = this->open_list.back();
    // Remove that node from the open_list.
    this->open_list.pop_back();
    // Return the pointer.
    return node;
}

// Create ConstructFinalPath method to return the final path found from your A* search.

std::vector<RouteModel::Node> RoutePlanner::ConstructFinalPath(RouteModel::Node *current_node) {
    distance = 0.0f;
    std::vector<RouteModel::Node> path_found;

    // This method should take the current (final) node as an argument and iteratively follow the 
    // chain of parents of nodes until the starting node is found.
    while (current_node) {
        path_found.push_back(*current_node);
        if (current_node->parent){
            // For each node in the chain, add the distance from the node to its parent to the distance variable.
            distance += current_node->distance(*current_node->parent); 
        } 
        current_node = current_node->parent;
    }
    // The returned vector should be in the correct order: the start node should be the first element
    // of the vector, the end node should be the last element.
    reverse(path_found.begin(), path_found.end());
    distance *= m_Model.MetricScale(); // Multiply the distance by the scale of the map to get meters.
    return path_found;

}

// Write the A* Search algorithm.

void RoutePlanner::AStarSearch() {
    start_node -> visited = true;
    open_list.push_back(start_node);
    
    RouteModel::Node* current_node = nullptr;

    while ( !open_list.empty() )
    {
        //Use the NextNode() method to sort the open_list and return the next node.
        current_node = this->NextNode();
        // When the search has reached the end_node, use the ConstructFinalPath method to return the final path that was found.
        if ( current_node->distance(*this->end_node) == 0 ) { 
            // Store the final path in the m_Model.path attribute before the method exits. This path will then be displayed on the map tile.
            m_Model.path = this->ConstructFinalPath(current_node);
            return; 
        }
        // Use the AddNeighbors method to add all of the neighbors of the current node to the open_list.
        AddNeighbors(current_node);
    }
}

