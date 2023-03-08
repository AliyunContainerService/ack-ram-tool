#include <vector>
#include <json/json.h>

#include <alibabacloud/core/ServiceResult.h>

class DescribeClustersResult : public AlibabaCloud::ServiceResult
{

public:
    DescribeClustersResult();
    explicit DescribeClustersResult(const std::string &payload);
//    ~DescribeClustersResult();

    struct ClusterItem
    {
        std::string clusterId;
        std::string name;
    };

    std::vector<ClusterItem> getClusters() const;

protected:
    void parse(const std::string &payload);
private:
    std::vector<ClusterItem> clusters_;
};

DescribeClustersResult::DescribeClustersResult(): ServiceResult()
{}

DescribeClustersResult::DescribeClustersResult(const std::string &payload) : ServiceResult()
{
    parse(payload);
}

void DescribeClustersResult::parse(const std::string &payload)
{
    Json::Reader reader;
    Json::Value value;
    reader.parse(payload, value);
//    setRequestId(value["RequestId"].asString());
    auto clusters = value;
    for (auto valueCluster : clusters) {
        ClusterItem item;
        if (!valueCluster["cluster_id"].isNull()) {
            item.clusterId = valueCluster["cluster_id"].asString();
        }
        if (!valueCluster["name"].isNull()) {
            item.name = valueCluster["name"].asString();
        }
        clusters_.push_back(item);
    }

}

 std::vector<DescribeClustersResult::ClusterItem> DescribeClustersResult::getClusters()const
{
    return clusters_;
}
